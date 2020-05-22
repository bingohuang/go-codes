package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		panic("first panic")
	}()

	defer func() {
		panic("second panic")
	}()

	defer func() {
		panic("third panic")
	}()

	panic("main panic")
}
