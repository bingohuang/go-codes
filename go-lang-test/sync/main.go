package main

import (
	"fmt"
	"sync"
)

func main() {
	//lock()
	Sync()
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
