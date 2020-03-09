package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	test4()
	time.Sleep(time.Second)
}

var once sync.Once

func doOnce() {
	fmt.Println("test")
}

func test1() {

	go once.Do(doOnce)
	go doOnce()
}

func test2() {

	go func() {
		once.Do(doOnce)
	}()
	go doOnce()
}

func doInit() {
	once.Do(doOnce)
}

func test3() {
	go doInit()
	doInit()
}

func test4() {
	go once.Do(doOnce)
	go func() {
		once.Do(doOnce)
	}()
}
