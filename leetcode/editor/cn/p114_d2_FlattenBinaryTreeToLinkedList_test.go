// github.com/bingohuang/go-codes
/**
题目: 114. 二叉树展开为链表
难度: 2
地址: https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO114 struct {
	in  *TreeNode
	out *TreeNode
}

func Test114(t *testing.T) {
	// add test cases
	treeNode_in1 := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				3,
				nil,
				nil,
			},
			&TreeNode{
				4,
				nil,
				nil,
			},
		},
		&TreeNode{
			5,
			nil,
			&TreeNode{
				6,
				nil,
				nil,
			},
		},
	}
	treeNode_out1 := &TreeNode{
		1,
		nil,
		&TreeNode{
			2,
			nil,
			&TreeNode{
				3,
				nil,
				&TreeNode{
					4,
					nil,
					&TreeNode{
						5,
						nil,
						&TreeNode{
							6,
							nil,
							nil,
						},
					},
				},
			},
		},
	}
	tc := map[string]IO114{
		"1": IO114{treeNode_in1, treeNode_out1},
	}

	for k, v := range tc {
		// algo func
		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput: %v\n", v.in)
		flatten(v.in)
		fmt.Printf("\toutput: %v\n", v.in)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(v.in, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, v.in)
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
func flatten(root *TreeNode) {
	// 20200721
	// 1. 暴力法:
	// Runtime:4 ms, faster than 65.02% of Go online submissions.
	// Memory Usage:2.9 MB, less than 80.00% of Go online submissions.
	var nodes []*TreeNode
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		nodes = append(nodes, node)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	for i := 1; i < len(nodes); i++ {
		nodes[i-1].Left = nil
		nodes[i-1].Right = nodes[i]
	}
	// 2. 原地转换法

}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个二叉树，原地将它展开为一个单链表。
//
//
//
// 例如，给定二叉树
//
//     1
//   / \
//  2   5
// / \   \
//3   4   6
//
// 将其展开为：
//
// 1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6
// Related Topics 树 深度优先搜索
