// github.com/bingohuang/go-codes
/**
题目: 23. 合并K个排序链表
难度: 3
地址: https://leetcode-cn.com/problems/merge-k-sorted-lists/
*/
package test

import (
	"container/heap"
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO23 struct {
	in  []*ListNode
	out *ListNode
}

func Test23(t *testing.T) {
	// add test cases
	link_node_1 := &ListNode{
		1,
		&ListNode{
			4,
			&ListNode{
				5,
				nil,
			},
		},
	}
	link_node_2 := &ListNode{
		1,
		&ListNode{
			3,
			&ListNode{
				4,
				nil,
			},
		},
	}
	link_node_3 := &ListNode{
		2,
		&ListNode{
			6,
			nil,
		},
	}
	link_out := &ListNode{
		1,
		&ListNode{
			1,
			&ListNode{
				2,
				&ListNode{
					3,
					&ListNode{
						4,
						&ListNode{
							4,
							&ListNode{
								5,
								&ListNode{
									6,
									nil,
								},
							},
						},
					},
				},
			},
		},
	}
	tc := map[string]IO23{
		"1": {[]*ListNode{link_node_1, link_node_2, link_node_3}, link_out},
	}

	for k, v := range tc {
		// algo func
		out := mergeKLists(v.in)

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
func mergeKLists(lists []*ListNode) *ListNode {
	// 20200423:
	// 解题思路一：暴力法
	// 用一个数组保存所有链表中的数，然后对这个数组进行排序，再从头到尾将数组遍历一遍，
	// 生成一个排好序的链表。假设每个链表的平均长度为 n，整体的时间复杂度就是 O(nk×log(nk))。

	// 解题思路二：最小堆法
	// 1. 把最小的 1 从所有的 k 个链表头里选出来之后，把 1 从链表里删掉。
	// 2. 下一个最小的数，还是从所有的 k 个链表头里选出来。
	// 3. 以此类推，每一轮都比较 k 个新的链表头的大小，得出最后的结果。
	if len(lists) == 0 {
		return nil
	}
	lh := make(ListHeap, 0)
	for i := range lists {
		if lists[i] != nil {
			lh = append(lh, lists[i])
		}
	}
	heap.Init(&lh)

	h := &ListNode{
		Val:  -1,
		Next: nil,
	}
	t := h
	for len(lh) > 0 {
		// 获取K个链表的最小头
		node := heap.Pop(&lh).(*ListNode)

		// 给新链表加节点
		t.Next = node
		t = t.Next

		// 将next节点放入最小堆中
		if node.Next != nil {
			heap.Push(&lh, node.Next)
		}
	}
	return h.Next

}

type ListHeap []*ListNode

func (h ListHeap) Len() int           { return len(h) }
func (h ListHeap) Less(i, j int) bool { return h[i].Val < h[j].Val } // 最小堆里的会议室按照会议的结束时间排序。
func (h ListHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ListHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
//
// 示例:
//
// 输入:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//输出: 1->1->2->3->4->4->5->6
// Related Topics 堆 链表 分治算法
