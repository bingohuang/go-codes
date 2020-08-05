// github.com/bingohuang/go-codes
/**
题目: 2. 两数相加
难度: 2
地址: https://leetcode-cn.com/problems/add-two-numbers/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO2 struct {
	in1 *ListNode
	in2 *ListNode
	out *ListNode
}

func Test2(t *testing.T) {
	// add test cases
	tc1_in1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	tc1_in2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	tc1_out := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}
	tc2_in1 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}
	tc2_in2 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	tc2_out := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	tc := map[string]IO2{
		"0": IO2{nil, nil, nil},
		"1": IO2{tc1_in1, tc1_in2, tc1_out},
		"2": IO2{tc2_in1, tc2_in2, tc2_out},
	}

	for k, v := range tc {
		// algo func
		out := addTwoNumbers(v.in1, v.in2)

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 20200805
	// https://leetcode-cn.com/problems/add-two-numbers/solution/hua-jie-suan-fa-2-liang-shu-xiang-jia-by-guanpengc/
	// Runtime:8 ms, faster than 92.88% of Go online submissions.
	// Memory Usage:5 MB, less than 88.95% of Go online submissions.
	pre := &ListNode{Val: 0}
	cur := pre
	carry := 0
	for l1 != nil || l2 != nil {
		x := 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		y := 0
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		carry = sum / 10
		sum = sum % 10
		cur.Next = &ListNode{Val: sum}

		cur = cur.Next

	}
	if carry == 1 {
		cur.Next = &ListNode{Val: carry}
	}
	return pre.Next
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例：
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
// Related Topics 链表 数学
// 👍 4701 👎 0
