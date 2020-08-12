// github.com/bingohuang/go-codes
/**
题目: 剑指 Offer 06. 从尾到头打印链表
难度: 1
地址: https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO剑指Offer06 struct {
	in  *ListNode
	out []int
}

func Test剑指Offer06(t *testing.T) {
	// add test cases
	tc := map[string]IO剑指Offer06{
		"0": IO剑指Offer06{nil, nil},
	}

	for k, v := range tc {
		// algo func
		out := reversePrint(v.in)

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
func reversePrint(head *ListNode) []int {
	// 20200811
	// 1、递归法
	// Runtime:4 ms, faster than 56.10% of Go online submissions.
	// Memory Usage:4.6 MB, less than 20.87% of Go online submissions.
	if head == nil {
		return nil
	}
	return append(reversePrint(head.Next), head.Val)
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
//
//
// 示例 1：
//
// 输入：head = [1,3,2]
//输出：[2,3,1]
//
//
//
// 限制：
//
// 0 <= 链表长度 <= 10000
// Related Topics 链表
// 👍 45 👎 0
