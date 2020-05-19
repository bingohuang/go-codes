package main

import "fmt"

func main() {
	//panic1()
	panic2()
}

func decodeIE(data []byte) {
	decode := make([]byte, 0, 0)

	decode = append(decode, data[0])
	fmt.Printf("first decode element is:%v\n", decode)

	decode = append(decode, data[1])
	fmt.Printf("second decode element is:%v\n", decode)
}

func panic1() {

	bytes := make([]byte, 1, 3)

	for i := 0; i < len(bytes); i++ {
		bytes[i] = byte(i + 2)
	}

	decodeIE(bytes)

}

var Tbl = []int{7, 8, 9, 10}

func panic2() {
	for i := 0; i < 3; i++ {
		Tbl = append(Tbl[:i], Tbl[i+1:]...)
		fmt.Println(i, len(Tbl), Tbl)
	}
}
