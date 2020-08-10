// github.com/bingohuang/go-codes
/**
题目: 19. 删除链表的倒数第N个节点
难度: 2
地址: https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO19 struct {
	in1 *ListNode
	in2 int
	out *ListNode
}

func Test19(t *testing.T) {
	// add test cases
	// 给定一个链表: 1->2->3->4->5, 和 n = 2.
	//当删除了倒数第二个节点后，链表变为 1->2->3->5.
	tc := map[string]IO19{
		//"0": {&ListNode{},0, &ListNode{}},
		//"1": {&ListNode{Val:1},1, nil},
		"2": {&ListNode{Val: 1, Next: &ListNode{Val: 2}}, 2, &ListNode{Val: 2}},
		"3": {&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		}, 2, &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		}},
	}

	for k, v := range tc {
		// algo func
		out := removeNthFromEnd(v.in1, v.in2)

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 20200810
	// 1、快慢指针
	// 如果不用pre节点，代码写的很丑而且很难考虑所有情况
	/*if n == 0 {
		return head
	}
	slow := head
	quick := head

	for i:=0;i < n;i++ {
		quick = quick.Next
	}
	if quick == nil {
		if n == 1 {
			return nil
		} else {
			slow.Next = slow.Next.Next
			return head
		}
	}
	for quick.Next != nil {
		quick  = quick.Next
		slow = slow.Next
	}
	// 找到 slow 对应的节点
	if slow.Next.Next != nil {
		slow.Next = slow.Next.Next
	} else {
		slow.Next = nil
	}

	return head*/
	// 设置一个pre节点，代码贼好写
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.2 MB,击败了100.00% 的Go用户
	pre := &ListNode{}
	pre.Next = head
	slow := pre
	fast := pre
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return pre.Next
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
// 示例：
//
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
//
//
// 说明：
//
// 给定的 n 保证是有效的。
//
// 进阶：
//
// 你能尝试使用一趟扫描实现吗？
// Related Topics 链表 双指针
// 👍 931 👎 0
