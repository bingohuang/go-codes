package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			a <- i
		}
		//1
		//close(a)
	}()

	//2.
	close(a)
	for i := range a {
		fmt.Println("Value i", i)
	}
}
