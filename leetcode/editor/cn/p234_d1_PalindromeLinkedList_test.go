// github.com/bingohuang/go-codes
/**
题目: 234. 回文链表
难度: 1
地址: https://leetcode-cn.com/problems/palindrome-linked-list/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO234 struct {
	in  *ListNode
	out bool
}

func Test234(t *testing.T) {
	// add test cases
	// 示例 1:
	// 输入: 1->2
	//输出: false
	//
	// 示例 2:
	// 输入: 1->2->2->1
	//输出: true
	tc := map[string]IO234{
		"0": IO234{nil, true},
		"1": IO234{&ListNode{1, &ListNode{2, nil}}, false},
		"2": IO234{&ListNode{1, &ListNode{2,
			&ListNode{2, &ListNode{1, nil}}}}, true},
	}

	for k, v := range tc {
		// algo func
		out := isPalindrome(v.in)

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
func isPalindrome(head *ListNode) bool {
	// 20200811
	// 1、复制到数组，再用双指针比较数组是否为回文
	// 复杂度: O(N)/O(N)
	// Runtime:16 ms, faster than 78.06% of Go online submissions.
	// Memory Usage:7.1 MB, less than 14.91% of Go online submissions.
	/*if head == nil {
		return true
	}
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	first := 0
	last := len(list) - 1
	for first < last {
		if list[first] != list[last] {
			return false
		}
		first ++
		last --
	}
	return true*/

	// 2、递归法
	// 复杂度: O(N)/O(N)
	// Runtime:20 ms, faster than 17.09% of Go online submissions.
	// Memory Usage:10 MB, less than 5.59% of Go online submissions.
	front := head
	var recursive func(curr *ListNode) bool
	recursive = func(curr *ListNode) bool {
		if curr != nil {
			if !recursive(curr.Next) {
				return false
			}
			if curr.Val != front.Val {
				return false
			}
			front = front.Next
		}
		return true
	}
	return recursive(head)
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//请判断一个链表是否为回文链表。
//
// 示例 1:
//
// 输入: 1->2
//输出: false
//
// 示例 2:
//
// 输入: 1->2->2->1
//输出: true
//
//
// 进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
// Related Topics 链表 双指针
// 👍 597 👎 0
