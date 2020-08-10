// github.com/bingohuang/go-codes
/**
题目: 108. 将有序数组转换为二叉搜索树
难度: 1
地址: https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO108 struct {
	in  []int
	out *TreeNode
}

func Test108(t *testing.T) {
	// add test cases
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
	tc := map[string]IO108{
		"0": {nil, nil},
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
	// 20200810
	// 1、递归法：类似于面试题 04.02 求高度最小的二叉树
	// 复杂度：O(N)/O(logN)
	// 执行耗时:4 ms,击败了84.84% 的Go用户
	// 内存消耗:4.4 MB,击败了79.79% 的Go用户
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	middle := len(nums) / 2
	root := &TreeNode{Val: nums[middle]}
	root.Left = sortedArrayToBST(nums[:middle])
	root.Right = sortedArrayToBST(nums[middle+1:])
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
//
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
//
// 示例:
//
// 给定有序数组: [-10,-3,0,5,9],
//
//一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
//
//      0
//     / \
//   -3   9
//   /   /
// -10  5
//
// Related Topics 树 深度优先搜索
// 👍 540 👎 0
