package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "tan"
	b := []byte(s)
	fmt.Println(b)
	sort.Slice(b, func(i int, j int) bool { return b[i] < b[j] })
	fmt.Println(b)
	fmt.Println(string(b))

	//ns := []byte{'1', 'B', 'C', 'a', 'g', 'M', '9'}
	ns := []byte{'t', 'a', 'n'}
	sort.Slice(ns, func(i int, j int) bool { return ns[i] < ns[j] })
	fmt.Println(string(ns))
}
