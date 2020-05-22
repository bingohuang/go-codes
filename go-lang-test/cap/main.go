package main

import "fmt"

func main() {
	a := make(chan int)
	fmt.Println(cap(a))

	b := [1]int{0}
	fmt.Println(cap(b))

	c := []int{0}
	fmt.Println(cap(c))

	//d := make(map[int]int)
	//fmt.Println(cap(d))
}
