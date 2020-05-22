package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{5, 4, 3}
	copy(s2, s1)
	fmt.Println("s2:", s2)

	s2 = []int{5, 4, 3}
	copy(s1, s2)
	fmt.Println("s1:", s1)
}
