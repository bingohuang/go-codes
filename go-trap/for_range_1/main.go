package main

import "fmt"

func main() {
	for _, s := range "Hi,中国" {
		fmt.Printf("0x%X\t", s)
	}
	fmt.Println()
	for _, b := range []byte("Hi,中国") {
		fmt.Printf("0x%X\t", b)
	}
	fmt.Println()
}
