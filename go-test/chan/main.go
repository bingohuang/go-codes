package main

import (
	"fmt"
	"time"
)

func main() {
	//chan1()
	//chan2()
	//chan3()
	chan4()
}

func chan1() {
	ch := make(chan int, 1)
	ch <- 2
	close(ch)
	if v, ok := <-ch; ok {
		fmt.Println(v)
	}

}

func chan2() {
	ch := make(chan int, 1)

	go func() {
		for i := 0; i < 3; i++ {
			v, _ := <-ch
			fmt.Println(v)
			close(ch)
		}
	}()

	for i := 0; i < 3; i++ {
		ch <- i
	}

}

func chan3() {
	ch := make(chan int, 1)

	go func() {
		for i := 0; i < 3; i++ {
			v, _ := <-ch
			fmt.Println(v)
		}
		close(ch)
	}()

	for i := 0; i < 3; i++ {
		ch <- i
	}
	time.Sleep(time.Second)
}

func chan4() {
	ch1 := make(chan int)
	fmt.Println(ch1)
	ch2 := make(<-chan int)
	fmt.Println(ch2)
	ch3 := make(chan<- int)
	fmt.Println(ch3)
	ch4 := make(chan<- chan int)
	fmt.Println(ch4)

}
