// github.com/bingohuang/go-codes
/**
题目: 160. 相交链表
难度: 1
地址: https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO160 struct {
	in1 *ListNode
	in2 *ListNode
	out *ListNode
}

func Test160(t *testing.T) {
	// add test cases
	tc1_out := &ListNode{Val: 2, Next: &ListNode{Val: 4}}
	tc1_in1 := &ListNode{0, &ListNode{9, &ListNode{1, tc1_out}}}
	tc1_in2 := &ListNode{3, tc1_out}
	tc := map[string]IO160{
		"0": {nil, nil, nil},
		"1": {tc1_in1, tc1_in2, tc1_out},
	}

	for k, v := range tc {
		// algo func
		out := getIntersectionNode(v.in1, v.in2)

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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 20200810
	// intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4]
	// 1、暴力法：遍历A链表中的每个节点，在B链表中是否存在

	// 2、hash法：遍历A链表，建立hashmap，再遍历B链表，看是否有对应相同的节点
	// 该用例出错：
	// [4,1,8,4,5]
	// [5,6,1,8,4,5]
	/*m := make(map[int]bool)
	for headA != nil {
		m[headA.Val] = true
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := m[headB.Val]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil*/
	// 3、双指针法
	// 执行耗时:56 ms,击败了14.88% 的Go用户
	// 内存消耗:9.1 MB,击败了39.22% 的Go用户
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != pB {
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}

		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}
	return pA
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//编写一个程序，找到两个单链表相交的起始节点。
//
// 如下面的两个链表：
//
//
//
// 在节点 c1 开始相交。
//
//
//
// 示例 1：
//
//
//
// 输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, s
//kipB = 3
//输出：Reference of the node with value = 8
//输入解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1
//,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
//
//
//
//
// 示例 2：
//
//
//
// 输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB =
// 1
//输出：Reference of the node with value = 2
//输入解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4
//]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
//
//
//
//
// 示例 3：
//
//
//
// 输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
//输出：null
//输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。由于这两个链表不相交，所以 intersectVal 必须为 0，而
// skipA 和 skipB 可以是任意值。
//解释：这两个链表不相交，因此返回 null。
//
//
//
//
// 注意：
//
//
// 如果两个链表没有交点，返回 null.
// 在返回结果后，两个链表仍须保持原有的结构。
// 可假定整个链表结构中没有循环。
// 程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
//
// Related Topics 链表
// 👍 768 👎 0
