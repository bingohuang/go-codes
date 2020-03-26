package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		//go f1(i, &wg)
		go f2(i, wg)
	}
	wg.Wait()
}

// [正确]一定要通过指针传值，不然进程会进入死锁状态
func f1(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	wg.Done()
}

// [错误]fatal error: all goroutines are asleep - deadlock!
func f2(i int, wg sync.WaitGroup) {
	fmt.Println(i)
	wg.Done()
}
