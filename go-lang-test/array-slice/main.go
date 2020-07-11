package main

import "fmt"

func main() {
	array_slice1()
	//array_slice2()
}

func array_slice1() {
	array := [...]string{"go", "c", "c++", "java"}
	slice1 := array[0:2]
	slice3 := append(slice1, "python")
	slice3 = append(slice3, "ruby")
	//slice3 = append(slice3, "js")
	fmt.Println(array, slice3)
}

func array_slice2() {
	slice := []int{0, 1, 2, 3, 4}
	slice1 := slice[0:2]
	slice1 = append(slice1, 10)
	fmt.Println(slice)
	fmt.Println(slice1)
}
