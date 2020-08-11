// github.com/bingohuang/go-codes
/**
题目: 148. 排序链表
难度: 2
地址: https://leetcode-cn.com/problems/sort-list/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO148 struct {
	in  *ListNode
	out *ListNode
}

func Test148(t *testing.T) {
	// add test cases
	// 示例 1:
	// 输入: 4->2->1->3
	// 输出: 1->2->3->4
	//
	// 示例 2:
	// 输入: -1->5->3->4->0
	// 输出: -1->0->3->4->5
	tc := map[string]IO148{
		//"0": IO148{nil, nil},
		"1": IO148{&ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}},
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}},
		"2": IO148{&ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{0, nil}}}}},
			&ListNode{-1, &ListNode{0, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}},
	}

	for k, v := range tc {
		// algo func
		out := sortList(v.in)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput: %v\n", v.in)
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
func sortList(head *ListNode) *ListNode {
	// 20200811
	// 1、归并排序
	// 壮壮: https://leetcode-cn.com/problems/sort-list/solution/go-gui-bing-pai-xu-by-yi-jie-diao-min/
	// Runtime:16 ms, faster than 45.58% of Go online submissions.
	// Memory Usage:5.9 MB, less than 22.31% of Go online submission
	if head == nil || head.Next == nil {
		return head
	}
	// 1. find middle
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. recursive
	r := sortList(slow.Next)
	slow.Next = nil
	l := sortList(head)

	// 3. merge
	tmphead := &ListNode{Val: 0}
	var mergeList func(head, l, r *ListNode) *ListNode
	mergeList = func(head, l, r *ListNode) *ListNode {
		p := head
		for l != nil && r != nil {
			if l.Val < r.Val {
				p.Next = l
				l = l.Next
			} else {
				p.Next = r
				r = r.Next
			}
			p = p.Next
		}
		if l == nil {
			p.Next = r
		}
		if r == nil {
			p.Next = l
		}
		return head.Next
	}
	return mergeList(tmphead, l, r)
	// Runtime:8 ms, faster than 98.81% of Go online submissions.
	// Memory Usage:5.9 MB, less than 22.31% of Go online submission
	// 为空或者只有一个元素
	/*if head == nil || head.Next == nil {
		return head
	}
	// 1. 找中点，二分，左右分别进行排序,快慢指针 fast比slow快一个是为了让slow.next可以成为中点，便于截断
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	r := sortList(slow.Next)
	// 注意要将链表截断
	slow.Next = nil
	l := sortList(head)
	// 2. 合并，将两个list合并为一个
	tmphead := &ListNode{Val: 0}
	var mergeList func(temphead, l, r *ListNode) *ListNode
	mergeList = func(temphead, l, r *ListNode) *ListNode {
		p := temphead
		for l != nil && r != nil {
		if l.Val < r.Val {
		p.Next = l
		l = l.Next
	} else {
		p.Next = r
		r = r.Next
	}
		p = p.Next
	}
		if l == nil {
		p.Next = r
	}
		if r == nil {
		p.Next = l
	}
		return temphead.Next
	}

	return mergeList(tmphead, l, r)*/
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
//
// 示例 1:
//
// 输入: 4->2->1->3
//输出: 1->2->3->4
//
//
// 示例 2:
//
// 输入: -1->5->3->4->0
//输出: -1->0->3->4->5
// Related Topics 排序 链表
// 👍 677 👎 0
