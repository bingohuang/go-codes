package main

import "fmt"

func main() {
	string2()
}

func string1() {
	s := "12345678"
	fmt.Println(s[len(s)-2:])
}

func string2() {
	s := "hello你好"
	r := []rune(s)

	fmt.Println(len(s), len(r)) // 7,11

}
