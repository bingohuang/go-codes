package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	test3()
}

func Add(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func test1() {
	f := Add(1)

	//fmt.Println(f)

	f2 := Add(1)

	//fmt.Println(f)

	a := f()  //
	b := f2() //
	c := f()
	d := f()

	fmt.Println(a, b, c, d)
}

func test2() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Print(i, ",")
		}(i)
	}
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Print(i, ",")
		}()
	}

	time.Sleep(5 * time.Second)
}

func test3() {
	a := make(chan int)
	go func() {
		for i := 0; i < 2; i++ {
			time.Sleep(time.Second)
			a <- i
		}
		//close(a)
	}()

	close(a)
	for i := range a {
		fmt.Println("Value i", i)
	}
}
