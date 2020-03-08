package test

import "fmt"

/**
定义单链表结构
*/
type ListNode struct {
	Val  int
	Next *ListNode
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