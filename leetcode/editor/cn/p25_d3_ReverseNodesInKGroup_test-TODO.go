// github.com/bingohuang/go-codes
/**
题目: [25]K 个一组翻转链表
难度: 3
地址: https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
*/
package test

import (
	"reflect"
	"testing"
)

// input and ouput
type IO25 struct {
	in1 *ListNode
	in2 int
	out *ListNode
}

func Test25(t *testing.T) {
	// add test cases
	tc := map[string]IO25{
		"1": IO25{&ListNode{}, 2, &ListNode{}},
	}

	for k, v := range tc {
		// algo func
		out := reverseKGroup(v.in1, v.in2)
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
func reverseKGroup(head *ListNode, k int) *ListNode {

	return head
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 示例 :
//
// 给定这个链表：1->2->3->4->5
//
// 当 k = 2 时，应当返回: 2->1->4->3->5
//
// 当 k = 3 时，应当返回: 3->2->1->4->5
//
// 说明 :
//
//
// 你的算法只能使用常数的额外空间。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
// Related Topics 链表
