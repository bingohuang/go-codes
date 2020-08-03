// github.com/bingohuang/go-codes
/**
题目: 103. 二叉树的锯齿形层次遍历
难度: 2
地址: https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO103 struct {
	in  *TreeNode
	out [][]int
}

func Test103(t *testing.T) {
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
		"1": {tc1_in, [][]int{{3}, {20, 9}, {15, 7}}},
	}

	for k, v := range tc {
		// algo func
		out := zigzagLevelOrder(v.in)

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
func zigzagLevelOrder(root *TreeNode) [][]int {
	// 20200728
	// 1、bfs：层次遍历，按层次奇偶顺序反转
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.8 MB,击败了13.95% 的Go用户
	if root == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{root, nil}
	var level []int
	isOrderLeft := true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			if isOrderLeft {
				level = append(level, node.Val)
			} else {
				level = append([]int{node.Val}, level...)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		} else {
			res = append(res, level)
			level = make([]int, 0)
			if len(queue) > 0 {
				queue = append(queue, nil)
			}
			isOrderLeft = !isOrderLeft
		}
	}
	return res

	// 20200728
	// 2、dfs
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.8 MB,击败了33.33% 的Go用户
	/*if root == nil {
		return nil
	}
	var res [][]int
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if level >= len(res) {
			res = append(res, []int{node.Val})
		} else {
			if level % 2 == 0 {
				res[level] = append(res[level], node.Val)
			} else {
				res[level] = append([]int{node.Val}, res[level]...)
			}
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
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
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
// 返回锯齿形层次遍历如下：
//
// [
//  [3],
//  [20,9],
//  [15,7]
//]
//
// Related Topics 栈 树 广度优先搜索
// 👍 227 👎 0
