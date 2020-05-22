package main

import "fmt"

func main() {
	//map1()
	//map2()
	//map3()
	map4()
}

func map1() {
	m := make(map[int]string)
	m[1] = "h"
	m[2] = "w"
	fmt.Println(m)
	for k := range m {
		//fmt.Println(k)
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

func map4() {
	//var m map[string]int
	//m["one"] = 1
	//fmt.Println(m)

	m := make(map[int]int)
	m[1] = 1

	v1, ok := m[1]
	fmt.Println(v1, ok) // 1, true
	v2 := m[1]
	fmt.Println(v2) // 1
	v3, ok := m[2]
	fmt.Println(v3, ok) // 返回类型默认值和false
	v4 := m[2]
	fmt.Println(v4) // 返回类型默认值，这里不小心也很容易犯错
}
