// github.com/bingohuang/go-codes
/**
题目: 226. 翻转二叉树
难度: 1
地址: https://leetcode-cn.com/problems/invert-binary-tree/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO226 struct {
	in  *TreeNode
	out *TreeNode
}

func Test226(t *testing.T) {
	// add test cases
	tc1_in := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	tc1_out := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	tc := map[string]IO226{
		"1": {tc1_in, tc1_out},
	}

	for k, v := range tc {
		// algo func
		out := invertTree(v.in)

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
func invertTree(root *TreeNode) *TreeNode {
	// 20200722

	// 1、利用先序遍历
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.1 MB,击败了100.00% 的Go用户
	//if root == nil {
	//	return nil
	//}
	//right := root.Right
	//root.Right = invertTree(root.Left)
	//root.Left = invertTree(right)
	//return root

	// 2、利用中序遍历
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.1 MB,击败了88.89% 的Go用户
	//if root == nil {
	//	return nil
	//}
	//invertTree(root.Left)
	//right := root.Right
	//root.Right = root.Left
	//root.Left = right
	//invertTree(root.Left)
	//return root

	// 3、利用后续遍历
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.1 MB,击败了88.89% 的Go用户
	//if root == nil {
	//	return nil
	//}
	//left := invertTree(root.Left)
	//right := invertTree(root.Right)
	//root.Right = left
	//root.Left = right
	//return root

	// 4. 利用层次遍历
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.1 MB,击败了11.11% 的Go用户
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		right := n.Right
		n.Right = n.Left
		n.Left = right
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//翻转一棵二叉树。
//
// 示例：
//
// 输入：
//
//      4
//   /   \
//  2     7
// / \   / \
//1   3 6   9
//
// 输出：
//
//      4
//   /   \
//  7     2
// / \   / \
//9   6 3   1
//
// 备注:
//这个问题是受到 Max Howell 的 原问题 启发的 ：
//
// 谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
// Related Topics 树
// 👍 503 👎 0
