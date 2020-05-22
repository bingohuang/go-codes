package main

import "fmt"

func main() {
	defer1()
	//defer2()
}

func defer1() {
	if true {
		defer fmt.Println("1")
	} else {
		defer fmt.Println("2")
	}
	fmt.Println("3")
	// 3 1
}

func defer2() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i) // 3210
	}

}
