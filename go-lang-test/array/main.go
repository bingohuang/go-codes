package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//array1()
	array2()
}

func array1() {
	arr := [3]string{1: "Chris", 2: "Ron"}

	arr2 := [...]int{3: -1}

	fmt.Println(arr, arr2)
}

func array2() {
	var intValue = [4]int{1, 2, 3, 4}
	fmt.Println(intValue)

	unsafePtr := unsafe.Pointer((&intValue[0]))
	uintPtr := uintptr(unsafePtr) + 2*unsafe.Sizeof(intValue[0]) // right
	//uintPtr := uintptr(unsafePtr) + 2 // wrong
	*(*int)(unsafe.Pointer(uintPtr)) = 6

	//*(*int)(unsafePtr + 2) = 6 //  wrong
	//(*int)(unsafePtr + 2*unsafe.Sizeof(intValue[0])) = 6 //  wrong
	fmt.Println(intValue)
}
