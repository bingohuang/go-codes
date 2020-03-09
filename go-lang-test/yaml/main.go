package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Jeffail/gabs"
	jy "github.com/ghodss/yaml"
)

func main() {
	// yaml
	//buffer, err := ioutil.ReadFile("./test-deploy.yaml")
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//t := map[string]interface{}{}
	//err = yaml.Unmarshal(buffer, &t)
	//fmt.Printf("%v\n",t)

	//fi := "./microservices-frontend.yaml"
	fi := "./microservices.yaml"
	file, err := os.Open(fi)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fo, err := os.Create("./microservices-out.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()
	w := bufio.NewWriter(fo)

	scanner := bufio.NewScanner(file)
	yaml := ""
	for scanner.Scan() {
		line := scanner.Text()
		// # 开头代表注释
		if strings.HasPrefix(line, "#") {
			w.Write([]byte(line + "\n"))
			continue
		}
		// --- 开头，表示一段结束
		if strings.HasPrefix(line, "---") {
			fedYaml, _ := transfterToKubefed([]byte(yaml))
			w.Write(fedYaml)
			w.Write([]byte(line + "\n"))
			yaml = ""
			continue
		}
		yaml += line + "\n"
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	if err = w.Flush(); err != nil {
		panic(err)
	}
}

func transfterToKubefed(yaml []byte) ([]byte, error) {

	// yaml to json
	json, err := jy.YAMLToJSON(yaml)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	//fmt.Println(string(json))

	// gabs
	jsonParsed, _ := gabs.ParseJSON(json)

	kind := jsonParsed.Path("kind").Data().(string)

	name := jsonParsed.Path("metadata.name").Data().(string)
	//fmt.Println("name ==", name)

	metadata := jsonParsed.Path("metadata").Data()
	//fmt.Println("metadata ==", metadata)

	spec := jsonParsed.Path("spec").Data()
	//fmt.Println("spec ==", spec)

	jsonObj := gabs.New()
	jsonObj.Set("types.kubefed.io/v1beta1", "apiVersion")
	if kind == "Deployment" {
		jsonObj.Set("FederatedDeployment", "kind")
	} else if kind == "Service" {
		jsonObj.Set("FederatedService", "kind")
	} else {
		return yaml, fmt.Errorf("no such kind, return source")
	}
	jsonObj.Set(name, "metadata", "name")
	jsonObj.Set("micro-demo", "metadata", "namespace")
	jsonObj.Set(metadata, "spec", "template", "metadata")
	jsonObj.Set(spec, "spec", "template", "spec")
	cluster1 := make(map[string]string)
	cluster1["name"] = "cluster1"
	jsonObj.ArrayAppendP(cluster1, "spec.placement.clusters")
	cluster2 := make(map[string]string)
	cluster2["name"] = "cluster2"
	jsonObj.ArrayAppendP(cluster2, "spec.placement.clusters")

	//fmt.Println(jsonObj.String())

	// jsontoyaml
	result, err := jy.JSONToYAML(jsonObj.Bytes())
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	fmt.Println(string(result))
	return result, nil
}
