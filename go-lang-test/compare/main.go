package main

import "fmt"

func compareA(num uint16, bits uint8) bool {
	if num >= (1 << bits) {
		return true
	}
	return false
}

func compareB(num uint16, bits uint8) bool {
	if uint32(num) >= uint32(1<<bits) {
		return true
	}
	return false
}

func main() {
	fmt.Println(compareA(65535, 16))
	fmt.Println(compareB(65535, 16))
}
