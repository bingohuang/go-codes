package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	pkgInfo *build.Package
)
var (
	typeNames = flag.String("type", "", "必填，逗号连接的多个Type名")
)

func main() {
	flag.Parse()
	if len(*typeNames) == 0 {
		log.Fatal("-type 必填")
	}
	consts := getConsts()
	src := genString(consts)
	//保存到文件
	outputName := ""
	if outputName == "" {
		types := strings.Split(*typeNames, ",")
		baseName := fmt.Sprintf("%s_string.go", types[0])
		outputName = filepath.Join(".", strings.ToLower(baseName))
	}
	err := ioutil.WriteFile(outputName, src, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}
func getConsts() map[string][]string {
	//获得待处理的Type
	types := strings.Split(*typeNames, ",")
	typesMap := make(map[string][]string, len(types))
	for _, v := range types {
		typesMap[strings.TrimSpace(v)] = []string{}
	}
	//解析当前目录下包信息
	var err error
	pkgInfo, err = build.ImportDir(".", 0)
	if err != nil {
		log.Fatal(err)
	}
	fset := token.NewFileSet()
	for _, file := range pkgInfo.GoFiles {
		f, err := parser.ParseFile(fset, file, nil, 0)
		if err != nil {
			log.Fatal(err)
		}
		typ := ""
		ast.Inspect(f, func(n ast.Node) bool {
			decl, ok := n.(*ast.GenDecl)
			// 只需要const
			if !ok || decl.Tok != token.CONST {
				return true
			}
			for _, spec := range decl.Specs {
				vspec := spec.(*ast.ValueSpec)
				if vspec.Type == nil && len(vspec.Values) > 0 {
					// 排除 v = 1 这种结构
					typ = ""
					continue
				}
				//如果Type不为空，则确认typ
				if vspec.Type != nil {
					ident, ok := vspec.Type.(*ast.Ident)
					if !ok {
						continue
					}
					typ = ident.Name
				}
				//typ是否是需处理的类型
				consts, ok := typesMap[typ]
				if !ok {
					continue
				}
				//将所有const变量名保存
				for _, n := range vspec.Names {
					consts = append(consts, n.Name)
				}
				typesMap[typ] = consts
			}
			return true
		})
	}
	return typesMap
}
func genString(types map[string][]string) []byte {
	const strTmp = `
	package {{.pkg}}
	import "fmt"
	
	{{range $typ,$consts :=.types}}
	func (c {{$typ}}) String() string{
		switch c { {{range $consts}}
			case {{.}}:return "{{.}}"{{end}}
		}
		return fmt.Sprintf("Status(%d)", c)	
	}
	{{end}}
	`
	pkgName := os.Getenv("GOPACKAGE")
	if pkgName == "" {
		pkgName = pkgInfo.Name
	}
	data := map[string]interface{}{
		"pkg":   pkgName,
		"types": types,
	}
	//利用模板库，生成代码文件
	t, err := template.New("").Parse(strTmp)
	if err != nil {
		log.Fatal(err)
	}
	buff := bytes.NewBufferString("")
	err = t.Execute(buff, data)
	if err != nil {
		log.Fatal(err)
	}
	//格式化
	src, err := format.Source(buff.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	return src
}
