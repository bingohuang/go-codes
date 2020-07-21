// github.com/bingohuang/go-codes
/**
题目: 101. 对称二叉树
难度: 1
地址: https://leetcode-cn.com/problems/symmetric-tree/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO101 struct {
	in  *TreeNode
	out bool
}

func Test101(t *testing.T) {
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
	tc2_in := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	tc := map[string]IO101{
		"1": {tc1_in, true},
		"2": {tc2_in, false},
	}

	for k, v := range tc {
		// algo func
		out := isSymmetric(v.in)

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
func isSymmetric(root *TreeNode) bool {
	// 20200721
	// 1、递归：同时做先序遍历，一个从左往右，一个从右往左，对比是否相同 - 深度优先搜索
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.9 MB,击败了100.00% 的Go用户
	//var preorder func(p *TreeNode, q *TreeNode) bool
	//preorder = func(p *TreeNode, q *TreeNode) bool {
	//	if p == nil && q == nil {
	//		return true
	//	}
	//	if p == nil || q == nil {
	//		return false
	//	}
	//	return p.Val == q.Val && preorder(p.Left, q.Right) && preorder(p.Right, q.Left)
	//}
	//return preorder(root, root)

	// 20200721
	// 2、迭代：一个队列，同时放两个值取两个值 - 广度优先搜索
	// 执行耗时:4 ms,击败了75.60% 的Go用户
	// 内存消耗:3.1 MB,击败了28.57% 的Go用户
	p, q := root, root
	qu := []*TreeNode{}
	qu = append(qu, p)
	qu = append(qu, q)
	for len(qu) > 0 {
		p, q = qu[0], qu[1]
		qu = qu[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		qu = append(qu, p.Left)
		qu = append(qu, q.Right)
		qu = append(qu, p.Right)
		qu = append(qu, q.Left)
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个二叉树，检查它是否是镜像对称的。
//
//
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//     1
//   / \
//  2   2
// / \ / \
//3  4 4  3
//
//
//
//
// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//     1
//   / \
//  2   2
//   \   \
//   3    3
//
//
//
//
// 进阶：
//
// 你可以运用递归和迭代两种方法解决这个问题吗？
// Related Topics 树 深度优先搜索 广度优先搜索
// 👍 914 👎 0
