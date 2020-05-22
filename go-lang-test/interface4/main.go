package main

import "fmt"

func main() {
	var data *byte
	var in interface{}
	fmt.Println(data == nil)
	fmt.Println(in == nil)
	in = data
	fmt.Println(in == nil)
}
