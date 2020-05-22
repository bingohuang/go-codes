package main

import "fmt"

func main() {
	var ch = make(chan int)
	go func() {
		fmt.Println("do work")
		ch <- 1
	}()

	ch <- 2
	close(ch)
	fmt.Println("fun exit")
}
