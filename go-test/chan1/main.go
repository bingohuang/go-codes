package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	var cc = make(chan int)
	go client(cc)

	for {
		select {
		case <-cc: // Bad：当channel cc被关闭后如果不做检查则造成死循环
			fmt.Println("continue")
		case <-time.After(5 * time.Second):
			fmt.Println("timeout")
		}
	}
}

func client(c chan int) {
	defer close(c)

	for {
		err := processBusiness()
		if err != nil {
			c <- 0
			return
		}
		c <- 1
	}
}

func processBusiness() error {
	return errors.New("domo")
}
