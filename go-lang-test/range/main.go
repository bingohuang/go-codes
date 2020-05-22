package main

import "fmt"

func main() {
	mp := make(map[string][]int)
	mp["a"] = []int{1, 2}
	mp["b"] = []int{3, 4}
	for _, v := range mp {
		for i := 0; i < len(v); i++ {
			v[i]++
		}
	}
	fmt.Println(mp["a"])
	for _, v := range mp["b"] {
		v++
	}
	fmt.Println(mp["b"])
}
