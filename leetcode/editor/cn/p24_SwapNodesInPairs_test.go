// github.com/bingohuang/go-codes
package test

import (
	"reflect"
	"testing"
)

// input and ouput
type IO24 struct {
	in  *ListNode
	out *ListNode
}

func Test24(t *testing.T) {
	// add test cases
	link_in1 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					nil,
				},
			},
		},
	}
	link_out1 := &ListNode{
		2,
		&ListNode{
			1,
			&ListNode{
				4,
				&ListNode{
					3,
					nil,
				},
			},
		},
	}
	link_in2 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				nil,
			},
		},
	}
	link_out2 := &ListNode{
		2,
		&ListNode{
			1,
			&ListNode{
				3,
				nil,
			},
		},
	}
	tc := map[string]IO24{
		"1": IO24{link_in1, link_out1},
		"2": IO24{link_in2, link_out2},
	}

	for k, v := range tc {
		// algo func
		out := swapPairs2(v.in)
		//fmt.Printf("out=%v\nv.out=%v\n", out, v.out)
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

// 算法1
func swapPairs1(head *ListNode) *ListNode {
	prev := head
	for {
		if prev == nil || prev.Next == nil {
			break
		}
		prev.Val, prev.Next.Val = prev.Next.Val, prev.Val
		prev = prev.Next.Next
	}
	return head
}

// 算法2
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := head.Next
	head.Next = swapPairs2(res.Next)
	res.Next = head
	return res
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := head.Next
	head.Next = swapPairs(res.Next)
	res.Next = head
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例:
//
// 给定 1->2->3->4, 你应该返回 2->1->4->3.
//
// Related Topics 链表
