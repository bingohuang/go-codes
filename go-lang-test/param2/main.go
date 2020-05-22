package main

import "fmt"

func variadicParam(a ...int) {
	fmt.Println(a)
}
func Foo(a ...int) {
	variadicParam(a) // a
	//variadicParam(a...) // a
}
func main() {
	s := []int{1, 2, 3, 4}
	variadicParam(s)        // b
	variadicParam(s...)     // c
	variadicParam(s[1:]...) // d
}
