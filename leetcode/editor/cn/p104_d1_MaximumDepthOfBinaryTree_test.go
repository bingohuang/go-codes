// github.com/bingohuang/go-codes
/**
题目: 104. 二叉树的最大深度
难度: 1
地址: https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO104 struct {
	in  *TreeNode
	out int
}

func Test104(t *testing.T) {
	// add test cases
	tc1_in := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	tc := map[string]IO104{
		"1": {tc1_in, 3},
	}

	for k, v := range tc {
		// algo func
		out := maxDepth(v.in)

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
func maxDepth(root *TreeNode) int {
	// 20200721
	// 1、 递归-深度遍历，每次往下遍历，层次+1
	// 执行耗时:4 ms,击败了92.45% 的Go用户
	// 内存消耗:4.4 MB,击败了66.67% 的Go用户
	var max int
	var preorder func(node *TreeNode, length int)
	preorder = func(node *TreeNode, length int) {
		if node == nil {
			return
		}
		length++
		if length > max {
			max = length
		}
		preorder(node.Left, length)
		preorder(node.Right, length)
		length--
	}
	preorder(root, 0)
	return max
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个二叉树，找出其最大深度。
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回它的最大深度 3 。
// Related Topics 树 深度优先搜索
// 👍 611 👎 0
