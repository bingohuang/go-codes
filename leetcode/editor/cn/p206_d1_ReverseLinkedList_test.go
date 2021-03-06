// github.com/bingohuang/go-codes
/**
题目: [206]反转链表
难度: 1
地址: https://leetcode-cn.com/problems/reverse-linked-list/
*/
package test

import (
	"testing"
)

// input and ouput
type IO206 struct {
	in  *ListNode
	out *ListNode
}

func Test206(t *testing.T) {
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
		4,
		&ListNode{
			3,
			&ListNode{
				2,
				&ListNode{
					1,
					nil,
				},
			},
		},
	}
	tc := map[string]IO206{
		"1": IO206{link_in1, link_out1},
	}

	for k, v := range tc {
		// algo func
		printList(v.in)
		out := reverseList1(v.in)
		printList(out)
		printList(v.out)
		if !compareList(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

// 算法1: 翻转链表
func reverseList1(head *ListNode) *ListNode {
	var newHead, next *ListNode

	for head != nil {
		// 1. 保存head next的副本
		next = head.Next
		// 2. 修改head next指向new head
		head.Next = newHead
		// 3. 调整两个头节点
		newHead = head
		head = next
	}

	return newHead
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var newHead, next *ListNode

	for head != nil {
		next = head.Next
		head.Next = newHead
		newHead = head
		head = next
	}

	return newHead
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//反转一个单链表。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//
// 进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
// Related Topics 链表
