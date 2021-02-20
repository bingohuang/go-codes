package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	for _, v := range a {
		wg.Add(1)
		//go func() {
		//	time.Sleep(time.Second)
		//	fmt.Println(v)
		//	wg.Done()
		//}()
		go func(n int) {
			time.Sleep(time.Second)
			fmt.Println(n)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
