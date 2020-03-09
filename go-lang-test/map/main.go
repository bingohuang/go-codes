package main

import "fmt"

func main() {
	//map1()
	//map2()
	map3()
}

func map1() {
	m := make(map[int]string)
	m[1] = "h"
	m[2] = "w"
	fmt.Println(m)
	for k := range m {
		fmt.Println(k)
		m[k] = "k"
	}
	fmt.Println(m)

}

func map2() {
	m := make(map[int]string)
	m[1] = "h"
	m[5] = "w"
	fmt.Println(m)
	for i := 0; i < len(m); i++ {
		m[i] = "k"
	}
	fmt.Println(m)

}

func map3() {
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		m[i] = i
	}
	for k := range m {
		delete(m, k)
	}
	m = nil
}
