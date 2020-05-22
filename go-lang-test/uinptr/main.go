package main

import (
	"fmt"
	"unsafe"
)

func main() {
	test1 := struct {
	}{}
	test2 := struct {
	}{}

	p1 := uintptr(unsafe.Pointer(&test1))
	p2 := uintptr(unsafe.Pointer(&test2))

	if p1 == p2 {
		fmt.Println("equal")
	} else {
		fmt.Println("note equal")
	}

	b := unsafe.Sizeof(test1)
	if b == 0 {
		fmt.Println("zero")
	} else {
		fmt.Println("note zero")
	}
}
