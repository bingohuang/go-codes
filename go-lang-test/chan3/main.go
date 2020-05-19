package main

import (
	"fmt"
	"time"
)

func main() {
	var cc = make(chan int, 1)
	cc <- 5
	close(cc)
	for {
		select {
		case <-cc:
			fmt.Println("continue")
		case <-time.After(5 * time.Second):
			fmt.Println("timeout")
			return
		}
	}
}
