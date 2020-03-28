package main

import (
	"fmt"
)

func main() {
	a()
	fmt.Println(c())
	fmt.Println(d())

}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return

}

func c() (i int) {
	defer func() { i++ }()
	return 1 // 2
}

func d() (i int) {
	defer func() { i++ }()
	return i // 1
}
