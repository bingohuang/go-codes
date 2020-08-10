// github.com/bingohuang/go-codes
/**
题目: 102. 二叉树的层序遍历
难度: 2
地址: https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO102 struct {
	in  *TreeNode
	out [][]int
}

func Test102(t *testing.T) {
	// add test cases
	// [3,9,20,null,null,15,7]
	tc1_in := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	tc := map[string]IO103{
		"1": {tc1_in, [][]int{{3}, {9, 20}, {15, 7}}},
	}

	for k, v := range tc {
		// algo func
		out := levelOrder(v.in)

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
func levelOrder(root *TreeNode) [][]int {
	// 20200810
	// 1、dfs：递归的做层次遍历
	// 复杂度：O(N)/O(N)
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:3 MB,击败了5.56% 的Go用户
	/*if root == nil {
		return nil
	}
	var res [][]int
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if level >= len(res) {
			res = append(res, []int{node.Val})
		} else {
			res[level] = append(res[level], node.Val)
		}
		if node.Left != nil {
			dfs(node.Left, level + 1)
		}
		if node.Right != nil {
			dfs(node.Right, level + 1)
		}
	}
	dfs(root, 0)
	return res*/

	// 2、bfs：标准的层次遍历
	// 复杂度：O(N)/O(N)
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.8 MB,击败了90.48% 的Go用户
	if root == nil {
		return nil
	}
	var res [][]int
	q := []*TreeNode{root}
	for len(q) > 0 {
		l := len(q)
		var list []int
		for i := 0; i < l; i++ {
			node := q[0]
			q = q[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, list)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例：
//二叉树：[3,9,20,null,null,15,7],
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其层次遍历结果：
//
// [
//  [3],
//  [9,20],
//  [15,7]
//]
//
// Related Topics 树 广度优先搜索
// 👍 591 👎 0
