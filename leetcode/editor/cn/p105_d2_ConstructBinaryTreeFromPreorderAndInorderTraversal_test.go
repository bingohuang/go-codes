// github.com/bingohuang/go-codes
/**
题目: 105. 从前序与中序遍历序列构造二叉树
难度: 2
地址: https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO105 struct {
	in1 []int
	in2 []int
	out *TreeNode
}

func Test105(t *testing.T) {
	// add test cases
	//前序遍历 preorder = [3,9,20,15,7]
	//中序遍历 inorder = [9,3,15,20,7]
	//
	// 返回如下的二叉树：
	//
	//     3
	//   / \
	//  9  20
	//    /  \
	//   15   7
	tc := map[string]IO105{
		"0": {[]int{}, []int{}, nil},
		"1": {[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7},
			&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}},
	}

	for k, v := range tc {
		// algo func
		out := buildTree(v.in1, v.in2)

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
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 20200810
	// 1、递归法：
	// 复杂度：O(N)/O(N)
	// 执行耗时:12 ms,击败了7.27% 的Go用户
	// 内存消耗:3.9 MB,击败了81.42% 的Go用户
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//根据一棵树的前序遍历与中序遍历构造二叉树。
//
// 注意:
//你可以假设树中没有重复的元素。
//
// 例如，给出
//
// 前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//
// 返回如下的二叉树：
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
// Related Topics 树 深度优先搜索 数组
// 👍 612 👎 0
