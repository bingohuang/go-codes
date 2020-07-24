// github.com/bingohuang/go-codes
/**
题目: 236. 二叉树的最近公共祖先
难度: 2
地址: https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO236 struct {
	in1 *TreeNode
	in2 *TreeNode
	in3 *TreeNode
	out *TreeNode
}

func Test236(t *testing.T) {
	// add test cases
	tc1A := &TreeNode{
		Val: 1,
	}
	tc1B := &TreeNode{
		Val: 2,
	}
	tc1C := &TreeNode{
		Val: 3,
	}
	tc1D := &TreeNode{
		Val: 4,
	}
	tc1B.Left = tc1A
	tc1B.Right = tc1C
	tc1C.Right = tc1D

	tc := map[string]IO236{
		//"1": {tc1B, tc1A,tc1C, tc1B},
		"2": {tc1B, tc1C, tc1D, tc1C},
	}

	for k, v := range tc {
		// algo func
		out := lowestCommonAncestor2(v.in1, v.in2, v.in3)

		fmt.Printf("case-%v:\n", k)
		//fmt.Printf("\tinput: %v\n", v.in)
		fmt.Printf("\toutput: %v\n", out)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
// 提交代码需修改函数名
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	// 20200724
	// 1、递归法：代码简单，但思路不是很好理解
	// 执行耗时:16 ms,击败了65.05% 的Go用户
	// 内存消耗:7.1 MB,击败了50.00% 的Go用户
	//if root == nil {
	//	return nil
	//}
	//if root.Val == p.Val || root.Val == q.Val {
	//	return root
	//}
	//left := lowestCommonAncestor(root.Left, p, q)
	//right := lowestCommonAncestor(root.Right, p, q)
	//if left != nil && right != nil {
	//	return root
	//}
	//if left == nil {
	//	return right
	//}
	//return left

	// 20200724
	// 2、迭代法：思路清晰，代码功底要好
	// 执行耗时:20 ms,击败了21.08% 的Go用户
	// 内存消耗:8.2 MB,击败了50.00% 的Go用户
	parent := map[int]*TreeNode{}
	visited := map[int]bool{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parent[node.Left.Val] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			parent[node.Right.Val] = node
			dfs(node.Right)
		}
	}
	dfs(root)

	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（
//一个节点也可以是它自己的祖先）。”
//
// 例如，给定如下二叉树: root = [3,5,1,6,2,0,8,null,null,7,4]
//
//
//
//
//
// 示例 1:
//
// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
//输出: 3
//解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
//
//
// 示例 2:
//
// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
//输出: 5
//解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
//
//
//
//
// 说明:
//
//
// 所有节点的值都是唯一的。
// p、q 为不同节点且均存在于给定的二叉树中。
//
// Related Topics 树
// 👍 666 👎 0
