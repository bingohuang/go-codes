// github.com/bingohuang/go-codes
/**
题目: 面试题 04.02. 最小高度树
难度: 1
地址: https://leetcode-cn.com/problems/minimum-height-tree-lcci/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO面试题0402 struct {
	in  []int
	out *TreeNode
}

func Test面试题0402(t *testing.T) {
	// add test cases
	// [0,-3,9,-10,null,5]
	tc1_out := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: -10,
			},
		},
		Right: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}
	tc := map[string]IO面试题0402{
		"1": {[]int{-10, -3, 0, 5, 9}, tc1_out},
	}

	for k, v := range tc {
		// algo func
		out := sortedArrayToBST(v.in)

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
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	// 20200727
	// 1、递归法
	// 执行耗时:8 ms,击败了9.61% 的Go用户
	// 内存消耗:4.4 MB,击败了100.00% 的Go用户
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	middle := len(nums) / 2
	head := &TreeNode{Val: nums[middle]}
	head.Left = sortedArrayToBST(nums[:middle])
	head.Right = sortedArrayToBST(nums[middle+1:])

	return head
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。
//示例: 给定有序数组: [-10,-3,0,5,9], 一个可能的答案是：[0,-3,9,-10,null,5]，
// 它可以表示下面这个高度平衡二叉搜索树：
//	  0
//	 / \
//  -3  9
//  /   /
// -10  5
//Related Topics 树 深度优先搜索
// 👍 42 👎 0
