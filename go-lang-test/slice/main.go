package main

import "fmt"

func main() {
	//slice1()
	//slice2()
	//slice3()
	slice4()
}

func slice1() {

	a := []int{1, 2, 3, 4, 5}
	b := a[3:4:4]

	fmt.Println(b)
	b[0] = -1
	fmt.Println(b)
	fmt.Println(a)
	b = append(b, 6)
	fmt.Println(b)
	fmt.Println(a)
}

func slice2() {

	a := []int{1, 2, 3, 4, 5}
	b := a[1:2:2]

	fmt.Println(b)
	b[0] = -1
	fmt.Println(b)
	fmt.Println(a)
	b = append(b, 6)
	fmt.Println(b)
	fmt.Println(a)
}

func slice3() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:3]
	fmt.Println("s2:", s2) // 2,3

	s1[1] = 0
	fmt.Println("s1:", s1) // 1,0,3
	fmt.Println("s2:", s2) // 0,3

	s2 = append(s2, 4)
	fmt.Println("s2:", s2) // 0,3,4
	s2[1] = 6
	fmt.Println("s2:", s2) // 2,6,4

	fmt.Println("s1:", s1) // 1,0,3
}

func slice4() {
	s := []int{1, 2}
	s2 := s            // 1 2
	s2 = append(s2, 3) // 1 2 3
	fmt.Println("s2:", len(s2), cap(s2))
	sliceAddOne(s)
	sliceAddOne(s2)
	fmt.Println(s, s2)
}

func sliceAddOne(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i]++
	}
}
