// github.com/bingohuang/go-codes
/**
题目: 剑指 Offer 26. 树的子结构
难度: 2
地址: https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO剑指Offer26 struct {
	in1 *TreeNode
	in2 *TreeNode
	out bool
}

func Test剑指Offer26(t *testing.T) {
	// add test cases
	tc1_in1 := &TreeNode{
		4,
		&TreeNode{2,
			&TreeNode{4,
				&TreeNode{Val: 8},
				&TreeNode{Val: 9},
			}, &TreeNode{Val: 5}},
		&TreeNode{3, &TreeNode{Val: 6}, &TreeNode{Val: 7}},
	}
	tc1_in2 := &TreeNode{4, &TreeNode{Val: 8}, &TreeNode{Val: 9}}
	tc := map[string]IO剑指Offer26{
		"0": IO剑指Offer26{nil, nil, false},
		"1": IO剑指Offer26{tc1_in1, tc1_in2, true},
	}

	for k, v := range tc {
		// algo func
		out := isSubStructure(v.in1, v.in2)

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
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	// 20200811
	// Runtime:24 ms, faster than 95.91% of Go online submissions.
	// Memory Usage:6.6 MB, less than 29.33% of Go online submissions.
	if B == nil || A == nil {
		return false
	}
	if A.Val == B.Val {
		//return containTree(A, B)
		if containTree(A, B) {
			return true
		}
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
func containTree(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return containTree(A.Left, B.Left) && containTree(A.Right, B.Right)
}

// https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/solution/go-qian-xu-bian-li-by-yi-jie-diao-min-2/
// 前序遍历A
/*func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	// 根 左右
	if A.Val == B.Val {
		if containTree(A, B) {
			return true
		}

	}
	l := isSubStructure(A.Left, B)
	r := isSubStructure(A.Right, B)
	return l || r
}

// A 中是否包含B（A,B同根）
func containTree(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	// B not nil, A nil  没匹配完A就nil了  false
	if A == nil {
		return false
	}

	if A.Val != B.Val {
		return false
	}
	l := containTree(A.Left, B.Left)
	r := containTree(A.Right, B.Right)
	return l && r
}*/
//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
//
// B是A的子结构， 即 A中有出现和B相同的结构和节点值。
//
// 例如:
//给定的树 A:
//
// 3
// / \
// 4 5
// / \
// 1 2
//给定的树 B：
//
// 4
// /
// 1
//返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。
//
// 示例 1：
//
// 输入：A = [1,2,3], B = [3,1]
//输出：false
//
//
// 示例 2：
//
// 输入：A = [3,4,5,1,2], B = [4,1]
//输出：true
//
// 限制：
//
// 0 <= 节点个数 <= 10000
// Related Topics 树
// 👍 93 👎 0
