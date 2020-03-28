package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	n int64
}

func main() {
	c := Counter{n: 1}
	r1 := c.Value1()
	fmt.Printf("c.value1=%d\n", r1)
	c.Increase(1)
	r2 := c.Value1()
	fmt.Printf("c.value1=%d\n", r2)

	r3 := c.Value2()
	fmt.Printf("c.value2=%d\n", r3)
	c.Increase(1)
	r4 := c.Value2()
	fmt.Printf("c.value2=%d\n", r4)
}

func (c *Counter) Increase(d int64) (r int64) {
	fmt.Printf("c1=%p\n", c)
	c.Lock()
	c.n += d
	r = c.n
	c.Unlock()
	return
}

// 此方法的实现是有问题的。当它被调用时， 一个Counter属主值将被复制。
func (c Counter) Value1() (r int64) {
	fmt.Printf("c2=%p\n", &c)
	c.Lock()
	r = c.n
	c.Unlock()
	return
}

//将接收参数改为指针类型
func (c *Counter) Value2() (r int64) {
	fmt.Printf("c2=%p\n", &c)
	c.Lock()
	r = c.n
	c.Unlock()
	return
}
