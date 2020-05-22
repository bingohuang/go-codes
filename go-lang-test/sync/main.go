package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//lock()
	Sync()
	//lock2()
}

func lock() {
	lock := sync.Mutex{}
	lock.Lock()
	fmt.Println("a")
	defer lock.Unlock()
	fmt.Println("b")
	lock.Lock()
	fmt.Println("c")
	defer lock.Unlock()
	fmt.Println("d")

}

var g sync.Once

func Sync() {
	f := func() {
		fmt.Print("a")
		Sync()
	}
	fmt.Print("b")
	g.Do(f)
	fmt.Print("c")
}

func lock2() {
	var lock sync.Mutex
	lock.Lock()

	go func() {
		lock.Lock()
		fmt.Println("b")
		lock.Unlock()
	}()

	time.Sleep(time.Second)
	fmt.Println("a")
	lock.Unlock()
	time.Sleep(time.Second)

}
