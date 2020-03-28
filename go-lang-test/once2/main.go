package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func initi() {
	fmt.Println("Bingo")
}

func main() {
	// 1 - No
	//fmt.Println("1")
	//go once.Do(initi)
	//go initi()

	// 2-
	//fmt.Println("2")
	//go func() {
	//    once.Do(initi)
	//}()
	//go initi()

	// 3-
	//fmt.Println("3")
	//go doInit()
	//doInit()

	// 4-
	fmt.Println("4")
	go once.Do(initi)
	go func() {
		once.Do(initi)
	}()

	time.Sleep(time.Second)
}

func doInit() {
	once.Do(initi)
}
