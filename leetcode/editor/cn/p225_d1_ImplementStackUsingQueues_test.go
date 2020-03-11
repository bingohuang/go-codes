// github.com/bingohuang/go-codes
/**
题目: 225. 用队列实现栈
难度: 1
地址: https://leetcode-cn.com/problems/implement-stack-using-queues/
*/
package test

import (
	"container/list"
	"fmt"
	"testing"
)

// input and ouput
type IO225 struct {
	in  []string
	out []int
}

func Test225(t *testing.T) {
	// add test cases
	obj := Constructor3()
	fmt.Printf("obj: %v\n", obj)
	obj.Push3(1)
	obj.Push3(2)
	fmt.Printf("obj: %v\n", obj)
	param_2 := obj.Pop3()
	fmt.Printf("param_2: %v\n", param_2)
	param_3 := obj.Top3()
	fmt.Printf("param_3: %v\n", param_3)
	param_4 := obj.Empty3()
	fmt.Printf("param_4: %v\n", param_4)
	obj.Pop3()
	param_5 := obj.Empty3()
	fmt.Printf("param_5: %v\n", param_5)
}

//leetcode submit region begin(Prohibit modification and deletion)
type MyNode struct {
	Val  int
	Next *MyNode
}
type MyStack struct {
	top *MyNode
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	node := &MyNode{Val: x, Next: this.top}
	this.top = node
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	x := this.top.Val
	this.top.Next, this.top = nil, this.top.Next
	return x
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.top.Val
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.top == nil
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)

// 算法1: 切片 slice
// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2 MB,击败了10.71% 的Go用户
type MyStack1 struct {
	slice []int
}

/** Initialize your data structure here. */
func Constructor1() MyStack1 {
	return MyStack1{}
}

/** Push element x onto stack. */
func (this *MyStack1) Push1(x int) {
	this.slice = append(this.slice, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack1) Pop1() int {
	x := this.slice[len(this.slice)-1]
	this.slice = this.slice[:len(this.slice)-1]
	return x
}

/** Get the top element. */
func (this *MyStack1) Top1() int {
	return this.slice[len(this.slice)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack1) Empty1() bool {
	return len(this.slice) == 0
}

// 算法2: 标准库 list
// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2 MB,击败了100.00% 的Go用户
type MyStack2 struct {
	*list.List
}

/** Initialize your data structure here. */
func Constructor2() MyStack2 {
	return MyStack2{list.New()}
}

/** Push element x onto stack. */
func (this *MyStack2) Push2(x int) {
	this.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack2) Pop2() int {
	return this.Remove(this.Back()).(int)
}

/** Get the top element. */
func (this *MyStack2) Top2() int {
	return this.Back().Value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack2) Empty2() bool {
	return this.Len() == 0
}

// 算法3: 自定义List
// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2 MB,击败了75.00% 的Go用户
type MyStack3 struct {
	top *MyNode
}

/** Initialize your data structure here. */
func Constructor3() MyStack3 {
	return MyStack3{}
}

/** Push element x onto stack. */
func (this *MyStack3) Push3(x int) {
	node := &MyNode{Val: x, Next: this.top}
	this.top = node
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack3) Pop3() int {
	x := this.top.Val
	this.top.Next, this.top = nil, this.top.Next
	return x
}

/** Get the top element. */
func (this *MyStack3) Top3() int {
	return this.top.Val
}

/** Returns whether the stack is empty. */
func (this *MyStack3) Empty3() bool {
	return this.top == nil
}

/* 题目详情 */
//使用队列实现栈的下列操作：
//
//
// push(x) -- 元素 x 入栈
// pop() -- 移除栈顶元素
// top() -- 获取栈顶元素
// empty() -- 返回栈是否为空
//
//
// 注意:
//
//
// 你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合
//法的。
// 你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
// 你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。
//
// Related Topics 栈 设计
