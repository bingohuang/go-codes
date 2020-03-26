package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	const Max = 100000
	const NumSenders = 1000
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)
	dataCh := make(chan int)
	/**stopCh是一个额外的信号通道。它的发送者为dataCh数据通道的接收者。
	  它的接收者为dataCh数据通道的发送者。**/
	stopCh := make(chan struct{})
	// 发送者
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}
	// 接收者
	go func() {
		defer wgReceivers.Done()
		for value := range dataCh {
			if value == Max-1 {
				/** 此接收者同时也是stopCh通道的唯一发送者。尽管它不能安全地关闭dataCh数据通道，但它可以安全地关闭stopCh通道。**/
				close(stopCh)
				return
			}
			fmt.Println(value)
		}
	}()
	wgReceivers.Wait()
}
