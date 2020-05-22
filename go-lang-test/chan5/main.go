package main

import "fmt"

func main() {
    chan1 := make(chan int)
    select {
    case chan1 <- 1:
        fmt.Println("C 1")
    case <- chan1:
        fmt.Println("C 2")
    }

    var a int chan

    var a1 chan int
    var a2 <-chan int
    var a chan<- int
}
