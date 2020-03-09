package main

import "fmt"

func main() {
	defer2()

}

func defer1() {
	if true {
		defer fmt.Println("1")
	} else {
		defer fmt.Println("2")
	}
	fmt.Println("3")

}

func defer2() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i) // 3210
	}

}
