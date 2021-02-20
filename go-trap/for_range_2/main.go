package main

import (
	"fmt"
	"sort"
)

func main() {
	heros := map[int]string{
		1: "superman",
		2: "batman",
		3: "spiderman",
		4: "the flash",
	}
	for k, v := range heros {
		fmt.Printf("%v:%s\t", k, v)
	}
	fmt.Println()

	// 一种解决 for-range 随机的问题
	var indexes []int
	for k, _ := range heros {
		indexes = append(indexes, k)
	}
	sort.Ints(indexes)
	for _, idx := range indexes {
		fmt.Printf("%v:%s\t", idx, heros[idx])
	}
}
