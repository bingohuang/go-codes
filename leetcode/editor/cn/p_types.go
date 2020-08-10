package test

import "fmt"

/**
定义单链表结构
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/**
打印单链表
*/
func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d->", l.Val)
		l = l.Next
	}
	fmt.Println()
}

/**
单链表转化为字符串
*/
func toStringList(l *ListNode) string {
	var str string
	for l != nil {
		str += fmt.Sprintf("%d->", l.Val)
		l = l.Next
	}
	return str
}

/**
比较单链表
*/
func compareList(a *ListNode, b *ListNode) bool {
	if a == nil || b == nil {
		if a == nil && b == nil {
			return true
		}
		return false
	}

	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return true
}

/**
比较大小
*/
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 备忘录Map
var memoMap = make(map[*TreeNode]int)
