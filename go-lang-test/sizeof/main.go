package main

import (
	"fmt"
	"unsafe"
)

func main() {
	sizeof1()
	sizeof2()
}

func sizeof1() {
	a := struct{}{}
	b := struct {
		FieldA float32
		FieldB string
	}{0, "bar"}
	c := struct {
		FieldA int8
		FieldB float32
		FieldC int64
		FieldD string
	}{0, 0, 0, "bar"}

	d := struct {
		FieldA int8
		FieldD string
		FieldB float32
		FieldC int64
	}{0, "bar", 0, 0}

	fmt.Printf("%d ", unsafe.Sizeof(a))
	fmt.Printf("%d ", unsafe.Sizeof(b))
	fmt.Printf("%d \n", unsafe.Sizeof(c))
	fmt.Printf("%d \n", unsafe.Sizeof(d))
}

func sizeof2() {
	a := "foo"
	b := "fooo"
	c := &b
	fmt.Printf("a: %T, %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("b: %T, %d\n", b, unsafe.Sizeof(b))
	fmt.Printf("c: %T, %d\n", c, unsafe.Sizeof(c))

	m := make(map[int]int)
	s := []int{1}
	fmt.Printf("m: %T, %d\n", m, unsafe.Sizeof(m))
	fmt.Printf("s: %T, %d\n", s, unsafe.Sizeof(s))

}
