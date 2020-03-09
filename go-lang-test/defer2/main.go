package main

import (
	_ "fmt"
)

func main() {
	a()
	fmt.Println(c())

}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return

}

func c() (i int) {
	defer func() { i++ }()
	return 1
}
