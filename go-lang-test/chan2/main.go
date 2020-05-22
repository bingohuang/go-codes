package main

import (
	"fmt"
	"time"
)

func main() {
	chan1()
	//chan2()
	//chan3()
}

// 执行第二个close会报错
func chan1() {
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

// 会死循环的输出continue
func chan2() {
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

// a = <- ch操作会发生阻塞，后面的goroutine不会启动，阻塞点不会被唤醒，runtime会报错。
func chan3() {
	ch := make(chan int)
	var a int
	a = <-ch
	println("read chan: ", a)

	go func() {
		println("write chan")
		ch <- 10
	}()
}
