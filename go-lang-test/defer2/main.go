package main

import (
	"fmt"
)

func main() {
	a()
	fmt.Println(c())
	fmt.Println(d())
	fmt.Println(deferDemo())
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return // 0
}

func c() (i int) {
	defer func() { i++ }()
	return 1 // 2
}

func d() (i int) {
	defer func() { i++ }()
	return i // 1
}

func deferDemo() (r int) {
	defer func() {
		r++
	}()
	return 0 // 1
}
