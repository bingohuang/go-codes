package main

import "fmt"

func main() {
	//var nilChan chan int
	//nilChan <- 5   // 阻塞 fatal error: all goroutines are asleep - deadlock!
	//n := <-nilChan // 阻塞 fatal error: all goroutines are asleep - deadlock!
	//fmt.Println(n)

	var closedChan = make(chan int)
	close(closedChan)
	m := <-closedChan
	fmt.Println(m)  // int类型的零值：0
	closedChan <- 5 // panic: send on closed channel
}
