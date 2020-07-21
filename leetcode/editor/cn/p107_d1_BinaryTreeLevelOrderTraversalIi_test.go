// github.com/bingohuang/go-codes
/**
题目: 107. 二叉树的层次遍历 II
难度: 1
地址: https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO107 struct {
	in  *TreeNode
	out [][]int
}

func Test107(t *testing.T) {
	// add test cases
	tc1_in := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	tc2_in := &TreeNode{
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
	tc := map[string]IO107{
		"1": IO107{tc1_in, [][]int{
			{2, 2},
			{1},
		}},
		"2": IO107{tc2_in, [][]int{
			{3, 4, 4, 3},
			{2, 2},
			{1},
		}},
	}

	for k, v := range tc {
		// algo func
		out := levelOrderBottom(v.in)

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
func levelOrderBottom(root *TreeNode) [][]int {
	// 20200721
	// 1、首先想到的是广度优先搜索-BFS
	// Runtime:4 ms, faster than 27.72% of Go online submissions.
	// Memory Usage:6.4 MB, less than 20.00% of Go online submissions.
	var result [][]int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)
		list := make([]int, 0)

		for i := 0; i < l; i++ {
			node := queue[i]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append([][]int{list}, result...) // 插入到result最前面
		queue = queue[l:]                         // 此处注意，是l不是1
	}

	return result

	// 20200721
	// 2、深度优先搜索也是可以的 - DFS

}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//
// 例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其自底向上的层次遍历为：
//
// [
//  [15,7],
//  [9,20],
//  [3]
//]
//
// Related Topics 树 广度优先搜索
// 👍 267 👎 0
