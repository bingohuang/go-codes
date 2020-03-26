package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

func main() {
	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 1000
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)
	dataCh := make(chan int)
	/**stopCh是一个额外的信号通道。它的发送者为中间调解者。它的接收者为dataCh数据通道的所有的发送者和接收者。**/
	stopCh := make(chan struct{})
	/**toStop是一个用来通知中间调解者让其关闭信号通道stopCh的第二个信号通道。此第二个信号通道的发送者为dataCh数据通道的所有的发送者和接收者，它的接收者为中间调解者。它必须为一个缓冲通道。**/
	toStop := make(chan string, 1)
	var stoppedBy string
	// 中间调解者
	go func() {
		stoppedBy = <-toStop
		close(stopCh)
	}()
	// 发送者
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				//if value == 0 {
				//    //使用了一个尝试发送操作来向中间调解者发送信号。
				//    select {
				//    case toStop <- "发送者#" + id:
				//    default:
				//    }
				//    return
				//}
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}
	// 接收者
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()
			for {
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						//使用一个尝试发送操作来向中间调解者发送信号。
						select {
						case toStop <- "接收者#" + id:
						default:
						}
						return
					}
					fmt.Println(id, ":", value)
				}
			}
		}(strconv.Itoa(i))
	}
	wgReceivers.Wait()
	fmt.Println("被" + stoppedBy + "终止了")
}
