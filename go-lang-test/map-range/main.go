package main

import "fmt"

func main() {
	range1() // 3 3 3 3
	range2() // 顺序随机
}

func range1() {
	s := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for k, v := range s {
		m[k] = &v
	}

	for _, v := range m {
		fmt.Println(*v)
	}
}

func range2() {
	tagsViews := map[string]int{
		"unix":   0,
		"python": 1,
		"go":     2,
		"golang": 3,
		"devops": 4,
		"gc":     5,
	}
	for key, views := range tagsViews {
		fmt.Println("There are", views, "views for", key)
	}
}
