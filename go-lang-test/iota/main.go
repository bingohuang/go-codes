package main

import "fmt"

func main() {
	fmt.Println(a, b, c, d, e, f)
}

const (
	a = iota
	b = iota
	c = iota
)

const (
	d, e, f = iota, iota, iota
)
