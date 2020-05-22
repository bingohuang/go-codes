package main

import "fmt"

func add(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func main() {
	f := add(1)
	f2 := add(1)
	a := f()
	b := f2()
	c := f()
	d := f()
	fmt.Println(a, b, c, d)
}
