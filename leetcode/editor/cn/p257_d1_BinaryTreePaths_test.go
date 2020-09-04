// github.com/bingohuang/go-codes
/**
题目: 257. 二叉树的所有路径
难度: 1
地址: https://leetcode-cn.com/problems/binary-tree-paths/
*/
package test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

// input and ouput
type IO257 struct {
	in  *TreeNode
	out []string
}

func Test257(t *testing.T) {
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
	tc2_in := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	tc := map[string]IO257{
		"0": {nil, nil},
		"1": {tc1_in, []string{"1->2", "1->3"}},
		"2": {tc2_in, []string{"1->2->5", "1->3"}},
	}

	for k, v := range tc {
		// algo func
		out := binaryTreePaths(v.in)

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
func binaryTreePaths(root *TreeNode) []string {
	// 20200725
	// 1、递归法
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.4 MB,击败了100.00% 的Go用户
	/*paths := make([]string, 0)

	var dfs func(*TreeNode, string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		path += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
		}
		path += "->"
		dfs(node.Left, path)
		dfs(node.Right, path)
	}
	dfs(root, "")
	return paths*/

	// 20200725
	// 1、迭代法
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.4 MB,击败了100.00% 的Go用户
	/*if root == nil {
		return nil
	}
	paths := make([]string, 0)
	nodeStack := []*TreeNode{root}
	pathStack := []string{strconv.Itoa(root.Val)}
	for len(nodeStack) > 0 {
		node := nodeStack[0]
		nodeStack = nodeStack[1:]
		path := pathStack[0]
		pathStack = pathStack[1:]
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
		}
		if node.Left != nil {
			nodeStack = append(nodeStack, node.Left)
			pathStack = append(pathStack, path+"->"+strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			nodeStack = append(nodeStack, node.Right)
			pathStack = append(pathStack, path+"->"+strconv.Itoa(node.Right.Val))
		}
	}
	return paths*/

	// 2020-09-04 16:38 @bingohuang 每日一题回顾
	// 算法：1、前序遍历、递归法
	// 复杂度：O(N^2)/O(N^2)
	// 效率：执行耗时:0 ms,击败了100.00% 的Go用户
	//			内存消耗:2.4 MB,击败了35.94% 的Go用户
	/*var res []string
	var preorder func(node *TreeNode, path string)
	preorder = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		path += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			res = append(res, path)
			return
		}
		path += "->"
		if node.Left != nil {
			preorder(node.Left, path)
		}
		if node.Right != nil {
			preorder(node.Right, path)
		}
	}
	preorder(root, "")
	return res*/

	// 2020-09-04 17:02 @bingohuang
	// 算法：2、迭代法
	// 复杂度：O(N^2)/O(N^2)
	// 效率：执行耗时:4 ms,击败了11.86% 的Go用户
	//			内存消耗:2.4 MB,击败了31.25% 的Go用户
	if root == nil {
		return nil
	}
	var res []string
	nodes := []*TreeNode{root}
	paths := []string{strconv.Itoa(root.Val)}
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		path := paths[0]
		paths = paths[1:]
		if node.Left == nil && node.Right == nil {
			res = append(res, path)
		}
		if node.Left != nil {
			nodes = append(nodes, node.Left)
			paths = append(paths, path+"->"+strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
			paths = append(paths, path+"->"+strconv.Itoa(node.Right.Val))
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个二叉树，返回所有从根节点到叶子节点的路径。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例:
//
// 输入:
//
//   1
// /   \
//2     3
// \
//  5
//
//输出: ["1->2->5", "1->3"]
//
//解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
// Related Topics 树 深度优先搜索
// 👍 299 👎 0
