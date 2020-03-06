package main

import (
	"fmt"
)

//go:generate echo hello
//go:generate go run main.go
//go:generate echo file=$GOFILE pkg=$GOPACKAGE
func main() {
	fmt.Printf("main func")
}
