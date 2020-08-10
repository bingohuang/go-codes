// github.com/bingohuang/go-codes
/**
题目: 94. 二叉树的中序遍历
难度: 2
地址: https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO94 struct {
	in  *TreeNode
	out []int
}

func Test94(t *testing.T) {
	// add test cases
	// 输入: [1,null,2,3]
	//   1
	//    \
	//     2
	//    /
	//   3
	//
	//输出: [1,3,2]
	tc := map[string]IO94{
		"0": {nil, nil},
		"1": {&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, []int{1, 3, 2}},
	}

	for k, v := range tc {
		// algo func
		out := inorderTraversal(v.in)

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
func inorderTraversal(root *TreeNode) []int {
	// 20200810
	// 1、递归法
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2 MB,击败了65.59% 的Go用户
	// 时间复杂度：O(n)。递归函数 T(n)=2⋅T(n/2)+1。
	// 空间复杂度：最坏情况下需要空间O(n)，平均情况为O(logn)。
	/*var res []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res*/

	// 20200810
	// 2、迭代法：
	// 2.1、以下为队列深度遍历法，也是前序遍历算法
	/*if root == nil {
		return nil
	}
	var res []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res*/
	// 2.2、利用栈来实现中序迭代遍历
	// 时间复杂度：O(n)
	// 空间复杂度：O(n)
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2 MB,击败了65.59% 的Go用户
	var res []int
	var stack []*TreeNode
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			// push
			stack = append(stack, curr)
			curr = curr.Left
		}
		// pop
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)
		curr = curr.Right
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个二叉树，返回它的中序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [1,3,2]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 哈希表
// 👍 617 👎 0
