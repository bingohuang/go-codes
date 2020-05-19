package main

import "fmt"

func main() {
	intChan := make(chan int)
	stringChan := make(chan string)
	go func() {
		intChan <- 1
	}()
	go func() {
		stringChan <- "hello"
	}()

	select {
	case value := <-intChan:
		fmt.Println(value)
	case value := <-stringChan:
		fmt.Println(value)
	default:
		fmt.Println("exception")
	}
}
