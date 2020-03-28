package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			a <- i
		}
		close(a)
	}()

	//close(a)
	for i := range a {
		fmt.Println("Value ", i)
	}
}
