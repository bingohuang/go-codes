# Go 实现

> [225. 用队列实现栈 - 简单](https://leetcode-cn.com/problems/implement-stack-using-queues/)
> [完整代码实现](https://github.com/bingohuang/go-codes/blob/master/leetcode/editor/cn/p225_d1_ImplementStackUsingQueues_test.go)

## 思考

## 算法1: 切片 slice

### Go 实现
```go
// 算法1: 切片 slice
// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2 MB,击败了10.71% 的Go用户
type MyStack struct {
	slice []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.slice = append(this.slice, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	x := this.slice[len(this.slice)-1]
	this.slice = this.slice[:len(this.slice)-1]
	return x
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.slice[len(this.slice)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.slice) == 0
}

```
### 复杂度分析
- 时间复杂度：$O(1)$

## 算法2: 标准库 list

### Go 实现
```go
// 算法2: 标准库 list
// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2 MB,击败了100.00% 的Go用户
type MyStack struct {
	*list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{list.New()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.Remove(this.Back()).(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Back().Value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Len() == 0
}
```
### 复杂度分析
- 时间复杂度：$O(1)$

## 算法3: 自定义List

### Go 实现
```go
// 算法3: 自定义List
// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2 MB,击败了75.00% 的Go用户
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
```
### 复杂度分析
- 时间复杂度：$O(1)$