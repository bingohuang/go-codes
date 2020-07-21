// github.com/bingohuang/go-codes
/**
题目: 21. 合并两个有序链表
难度: 1
地址: https://leetcode-cn.com/problems/merge-two-sorted-lists/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO21 struct {
	in1 *ListNode
	in2 *ListNode
	out *ListNode
}

func Test21(t *testing.T) {
	// add test cases
	tc1_in1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	tc1_in2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	tc1_out := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
	}
	tc := map[string]IO21{
		"1": {tc1_in1, tc1_in2, tc1_out},
	}

	for k, v := range tc {
		// algo func
		out := mergeTwoLists(v.in1, v.in2)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput1: %v\n", v.in1)
		fmt.Printf("\tinput2: %v\n", v.in2)
		fmt.Printf("\toutput: %v\n", out)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 20170721
	// 1、递归法
	// 执行耗时:4 ms,击败了61.10% 的Go用户
	// 内存消耗:2.6 MB,击败了63.64% 的Go用户
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例：
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
// Related Topics 链表
// 👍 1159 👎 0
