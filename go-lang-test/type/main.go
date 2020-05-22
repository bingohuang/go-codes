package main

import "fmt"

type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(i int) {
	t.a = i
}

func main() {
	//t := T{}
	//T.Set(t, 2)
	//
	//g := &T{}
	//(*T).Set(g, 3)
	//
	//fmt.Println(T.Get(t))
	//fmt.Println((*T).Get(g)) //  invalid method expression T.Set (needs pointer receiver: (*T).Set)

	// OK
	//t := &T{}
	//t.Set(2)
	//fmt.Println(t.Get())

	// OK
	//t := T{}
	//t.Set(2)
	//fmt.Println(t.Get())

	// OK
	t := &T{}
	f1 := t.Set

	g := &T{}
	f2 := g.Set

	f1(2)
	f2(3)
	fmt.Println(t.Get())
	fmt.Println(g.Get())
}
