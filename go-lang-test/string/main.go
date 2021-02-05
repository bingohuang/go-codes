package main

import (
	"fmt"
	"strconv"
)

func main() {
	//string2()
	string3()
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

func string3() {
	s := "18+"
	r, _ := strconv.Atoi(s[0:2])
	fmt.Println(r)
}
