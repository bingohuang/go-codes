package main

import "fmt"

var v1 = 10
var v2 = '你'
var v3 = 17.8

func main() {
	fmt.Printf("%d %s %.1f\n", v1, string(v2), v3)
	fmt.Printf("%d %s %.1f\n", v1, v2, v3)
	fmt.Printf("%d %U %.1f\n", v1, v2, v3)
	fmt.Printf("%d %d %.1F\n", v1, v2, v3)
}
