// github.com/bingohuang/go-codes
/**
题目: 98. 验证二叉搜索树
难度: 2
地址: https://leetcode-cn.com/problems/validate-binary-search-tree/
*/
package test

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// input and ouput
type IO98 struct {
	in  *TreeNode
	out bool
}

func Test98(t *testing.T) {
	// add test cases
	// 20200810
	// 输入:
	//    2
	//   / \
	//  1   3
	//输出: true
	// 输入:
	//    5
	//   / \
	//  1   4
	//     / \
	//    3   6
	//输出: false
	tc := map[string]IO98{
		"0": {nil, true},
		"1": {&TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, true},
		"2": {&TreeNode{Val: 5, Left: &TreeNode{Val: 1},
			Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}, false},
		"3*": {&TreeNode{Val: 1, Left: &TreeNode{Val: 1}}, false},
	}

	for k, v := range tc {
		// algo func
		out := isValidBST(v.in)

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
func isValidBST(root *TreeNode) bool {
	// 20200810
	// 1、递归法
	// 时间复杂度：O(N)
	// 空间复杂度：O(N)
	// 执行耗时:8 ms,击败了77.07% 的Go用户
	// 内存消耗:5.5 MB,击败了49.79% 的Go用户
	/*var helper func(node *TreeNode, lower int, upper int) bool
	helper = func(node *TreeNode, lower int, upper int) bool{
		if node == nil {
			return true
		}
		if node.Val <= lower || node.Val >= upper {
			return false
		}
		return helper(node.Left, lower, node.Val) && helper(node.Right, node.Val, upper)

	}
	return helper(root, math.MinInt64, math.MaxInt64)*/
	// 2、迭代法：中序遍历，类似于94题，使用栈来判断，是否都是升序遍历
	// 时间复杂度：O(N)
	// 空间复杂度：O(N)
	// 执行耗时:8 ms,击败了77.07% 的Go用户
	// 内存消耗:5.5 MB,击败了46.86% 的Go用户
	var stack []*TreeNode
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if inorder >= root.Val {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征：
//
//
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
//
// 示例 1:
//
// 输入:
//    2
//   / \
//  1   3
//输出: true
//
//
// 示例 2:
//
// 输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
//
// Related Topics 树 深度优先搜索
// 👍 702 👎 0
