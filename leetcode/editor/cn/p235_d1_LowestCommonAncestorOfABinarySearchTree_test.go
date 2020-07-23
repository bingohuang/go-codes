// github.com/bingohuang/go-codes
/**
题目: 235. 二叉搜索树的最近公共祖先
难度: 1
地址: https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO235 struct {
	in1 *TreeNode
	in2 *TreeNode
	in3 *TreeNode
	out *TreeNode
}

func Test235(t *testing.T) {
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

	tc := map[string]IO235{
		//"1": {tc1B, tc1A,tc1C, tc1B},
		"2": {tc1B, tc1C, tc1D, tc1C},
	}

	for k, v := range tc {
		// algo func
		out := lowestCommonAncestor(v.in1, v.in2, v.in3)

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
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 20200723
	// 1、递归法
	// 执行耗时:16 ms,击败了99.03% 的Go用户
	// 内存消耗:6.8 MB,击败了66.67% 的Go用户
	//if root == nil {
	//	return nil
	//}
	//rVal, pVal, qVal := root.Val, p.Val, q.Val
	//if pVal > rVal && qVal > rVal {
	//	return lowestCommonAncestor(root.Right, p, q)
	//} else if pVal < rVal && qVal < rVal {
	//	return lowestCommonAncestor(root.Left, p, q)
	//} else {
	//	return root
	//}

	// 2、迭代法：
	// 执行耗时:24 ms,击败了81.49% 的Go用户
	// 内存消耗:6.8 MB,击败了66.67% 的Go用户
	rVal, pVal, qVal := root.Val, p.Val, q.Val
	for root != nil {
		rVal = root.Val
		if pVal > rVal && qVal > rVal {
			root = root.Right
		} else if pVal < rVal && qVal < rVal {
			root = root.Left
		} else {
			return root
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（
//一个节点也可以是它自己的祖先）。”
//
// 例如，给定如下二叉搜索树: root = [6,2,8,0,4,7,9,null,null,3,5]
//
//
//
//
//
// 示例 1:
//
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
//输出: 6
//解释: 节点 2 和节点 8 的最近公共祖先是 6。
//
//
// 示例 2:
//
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
//输出: 2
//解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
//
//
//
// 说明:
//
//
// 所有节点的值都是唯一的。
// p、q 为不同节点且均存在于给定的二叉搜索树中。
//
// Related Topics 树
// 👍 338 👎 0
