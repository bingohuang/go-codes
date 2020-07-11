package main

import "fmt"

func main() {
	a := new(struct{})
	b := new(struct{})
	println(a, b, a == b)

	c := new(struct{})
	d := new(struct{})
	fmt.Println(c, d, c == d)
}
