package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	k := make([]byte, 32)
	/** 【正确】使用rand.Read()生成安全随机数  **/
	if _, err := rand.Read(k); err != nil {
		fmt.Printf("rand.Read() error : %v \n", err)
	}
	fmt.Printf("rand.Read(): %v \n", k)
}
