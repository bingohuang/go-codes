package main

// 24. 两两交换链表中的节点
// https://leetcode-cn.com/problems/swap-nodes-in-pairs/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 递归法
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := head.Next
	head.Next = swapPairs(res.Next)
	res.Next = head
	return res
}
