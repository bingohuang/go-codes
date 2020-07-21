// github.com/bingohuang/go-codes
/**
题目: 113. 路径总和 II
难度: 2
地址: https://leetcode-cn.com/problems/path-sum-ii/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO113 struct {
	in1 *TreeNode
	in2 int
	out [][]int
}

func Test113(t *testing.T) {
	// add test cases
	treeNode_in1 := &TreeNode{
		5,
		&TreeNode{
			4,
			&TreeNode{
				11,
				&TreeNode{
					7,
					nil,
					nil,
				},
				&TreeNode{
					2,
					nil,
					nil,
				},
			},
			nil,
		},
		&TreeNode{
			8,
			&TreeNode{
				13,
				nil,
				nil,
			},
			&TreeNode{
				4,
				&TreeNode{
					5,
					nil,
					nil,
				},
				&TreeNode{
					1,
					nil,
					nil,
				},
			},
		},
	}
	tc := map[string]IO113{
		"1": IO113{treeNode_in1, 22, [][]int{
			{5, 4, 11, 2}, {5, 8, 4, 5},
		}},
		"2": IO113{&TreeNode{}, 1, nil},
	}

	for k, v := range tc {
		// algo func
		out := pathSum(v.in1, v.in2)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput: %v\n", v.in1)
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
func pathSum(root *TreeNode, sum int) [][]int {
	// 20200721
	// 执行耗时:4 ms,击败了92.24% 的Go用户
	// 内存消耗:4.5 MB,击败了83.33% 的Go用户
	var result [][]int
	var path []int
	var pathValue int

	var preorder func(node *TreeNode, sum int)
	preorder = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		pathValue += node.Val
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && sum == pathValue {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
		}
		preorder(node.Left, sum)
		preorder(node.Right, sum)
		pathValue -= node.Val
		path = path[:len(path)-1]

	}
	preorder(root, sum)

	return result
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例:
//给定如下二叉树，以及目标和 sum = 22，
//
//               5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \    / \
//        7    2  5   1
//
//
// 返回:
//
// [
//   [5,4,11,2],
//   [5,8,4,5]
//]
//
// Related Topics 树 深度优先搜索
func pathSum1(root *TreeNode, sum int) [][]int {
	// 20200311
	var result [][]int
	var path []int
	var pathValue int
	var preOrder func(node *TreeNode, sum int)
	preOrder = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		pathValue += node.Val
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && pathValue == sum {
			t := make([]int, len(path))
			copy(t, path)
			result = append(result, t)
		}
		preOrder(node.Left, sum)
		preOrder(node.Right, sum)
		pathValue -= node.Val
		path = path[:len(path)-1] // pop
	}
	preOrder(root, sum)
	return result
}
