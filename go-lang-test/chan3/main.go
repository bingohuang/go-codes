package main

import (
	"fmt"
	"time"
)

var global int = 0
var c = make(chan int, 1) // 实现互斥锁
//var c = make(chan int, 0)

// 实现互斥锁
func thread1() {
	c <- 1
	global++
	time.Sleep(time.Second)
	<-c
}

func thread2() {
	<-c
	global++
	time.Sleep(time.Second)
	c <- 1
}

func main() {
	go thread2()
	go thread2()
	fmt.Println("global=", global)
	fmt.Println("c=", c)

	go thread2()
	fmt.Println("global=", global)
	fmt.Println("c=", c)

	time.Sleep(3 * time.Second)
	fmt.Println("global=", global)
	fmt.Println("c=", c)
}
