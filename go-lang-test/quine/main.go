// Go自重写代码

//package main;func main(){c:="package main;func main(){c:=%q;print(c,c)}";print(c,c)}

package main

func main() { print(c + "\x60" + c + "\x60") }

var c = `package main;func main(){print(c+"\x60"+c+"\x60")};var c=`

//package main
//
//import "fmt"
//
//func main() {
//    fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
//}
//
//var q = `/* Go quine */
//package main
//
//import "fmt"
//
//func main() {
//    fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
//}
//
//var q = `
