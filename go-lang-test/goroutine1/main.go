package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//go1()
	//go2()
	go3()
	//go4()
}

func go1() {
	go func() {
		fmt.Println("Go Goroutine!")
	}()
	runtime.Gosched()
}

func go2() {
	go func() {
		fmt.Println("Go Goroutine!")
	}()

}

func go3() {
	go func() {
		fmt.Println("Go Goroutine!")
	}()

	time.Sleep(time.Nanosecond)

}

func go4() {
	go func() {
		fmt.Println("Go Goroutine!")
	}()
	runtime.Goexit()

}
