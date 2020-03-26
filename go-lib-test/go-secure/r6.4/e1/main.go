package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var x int32 = 0
	for i := 0; i < 100; i++ {
		wg.Add(1) // [正确]将对Add方法的调用移出匿名协程之外
		go func() {
			//wg.Add(1) // [错误]Add方法的调用位置是不合适
			atomic.AddInt32(&x, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(atomic.LoadInt32(&x))
}
