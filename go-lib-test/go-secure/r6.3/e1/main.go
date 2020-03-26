package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	const Max = 100000
	const NumReceivers = 100
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)
	dataCh := make(chan int)
	// 发送者
	go func() {
		for {
			if value := rand.Intn(Max); value == 0 {
				// 此唯一的发送者可以安全地关闭此数据通道。
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()
	// 接收者
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()
			// 接收数据直到通道dataCh已关闭并且dataCh的缓冲队列已空。
			for value := range dataCh {
				fmt.Println(value)
			}
		}()
	}
	wgReceivers.Wait()
}
