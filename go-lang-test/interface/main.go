package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println("hello")
}

func main() {
	var t *T
	var i I = t
	t.M()
	i.M()
}
