// github.com/bingohuang/go-codes
/**
题目: 100. 相同的树
难度: 1
地址: https://leetcode-cn.com/problems/same-tree/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO100 struct {
	in1 *TreeNode
	in2 *TreeNode
	out bool
}

func Test100(t *testing.T) {
	// add test cases
	tc1_in1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	tc1_in2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	tc2_in2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	tc := map[string]IO100{
		"1": {tc1_in1, tc1_in2, true},
		"2": {tc1_in1, tc2_in2, false},
	}

	for k, v := range tc {
		// algo func
		out := isSameTree(v.in1, v.in2)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput1: %v\n", v.in1)
		fmt.Printf("\tinput2: %v\n", v.in2)
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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 20200721
	// 1、反射比较法-这个方法可行
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.3 MB,击败了60.00% 的Go用户
	//return reflect.DeepEqual(p, q)

	// 20200721
	// 2、暴力法：先序遍历，再比较字符串 - 这个方法不行，分辨不出来
	//var pStr, qStr string
	//var preorder func(node *TreeNode, str string) string
	//preorder = func(node *TreeNode, str string) string {
	//	if node == nil {
	//		return str
	//	}
	//	str += strconv.Itoa(node.Val) + "-"
	//	str = preorder(node.Left, str)
	//	str = preorder(node.Right, str)
	//	return str
	//}
	//pStr = preorder(p, pStr)
	//qStr = preorder(q, qStr)
	//return pStr == qStr

	// 20200721
	// 3、递归法
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.1 MB,击败了100.00% 的Go用户
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

	// 4、迭代法：双队列 - TODO
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定两个二叉树，编写一个函数来检验它们是否相同。
//
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
//
// 示例 1:
//
// 输入:       1         1
//          / \       / \
//         2   3     2   3
//
//        [1,2,3],   [1,2,3]
//
//输出: true
//
// 示例 2:
//
// 输入:      1          1
//          /           \
//         2             2
//
//        [1,2],     [1,null,2]
//
//输出: false
//
//
// 示例 3:
//
// 输入:       1         1
//          / \       / \
//         2   1     1   2
//
//        [1,2,1],   [1,1,2]
//
//输出: false
//
// Related Topics 树 深度优先搜索
// 👍 400 👎 0
