package main

import "fmt"

func main() {
	var a = []interface{}{123, "abc"}

	Print(a...)
	Print(a)

	Print([]interface{}{123, "abc"}...)
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}

// 可变参数的定义
func sum(vals ...int) int {
	return 0
}
