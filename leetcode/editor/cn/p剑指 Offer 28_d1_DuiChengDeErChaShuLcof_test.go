// github.com/bingohuang/go-codes
/**
题目: 剑指 Offer 28. 对称的二叉树
难度: 1
地址: https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO剑指Offer28 struct {
	in  *TreeNode
	out bool
}

func Test剑指Offer28(t *testing.T) {
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
	tc := map[string]IO剑指Offer28{
		"1": {tc1_in, true},
		"2": {tc2_in, false},
	}

	for k, v := range tc {
		// algo func
		out := isSymmetric2(v.in)

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
func isSymmetric2(root *TreeNode) bool {
	// 20200728
	// 1、递归法
	// 执行用时： 4 ms , 在所有 Go 提交中击败了 75.43% 的用户
	// 内存消耗： 2.9 MB , 在所有 Go 提交中击败了 100.00% 的用户
	/*var dfs func(p *TreeNode, q *TreeNode) bool
	  dfs = func(p *TreeNode, q *TreeNode) bool {
	      if p == nil && q == nil {
	          return true
	      }
	      if p == nil || q == nil {
	          return false
	      }
	      return p.Val == q.Val && dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
	  }
	  return dfs(root, root)*/

	// 2、迭代法
	// 执行用时： 4 ms , 在所有 Go 提交中击败了 75.43% 的用户
	// 内存消耗： 3 MB , 在所有 Go 提交中击败了 100.00% 的用户
	qu := []*TreeNode{root, root}
	for len(qu) > 0 {
		p, q := qu[0], qu[1]
		qu = qu[2:]
		if p == nil && q == nil {
			continue // 此处不能return，必须continue
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
//请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
// 1
// / \
// 2 2
// / \ / \
//3 4 4 3
//但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
// 1
// / \
// 2 2
// \ \
// 3 3
//
//
//
// 示例 1：
//
// 输入：root = [1,2,2,3,4,4,3]
//输出：true
//
//
// 示例 2：
//
// 输入：root = [1,2,2,null,3,null,3]
//输出：false
//
//
//
// 限制：
//
// 0 <= 节点个数 <= 1000
//
// 注意：本题与主站 101 题相同：https://leetcode-cn.com/problems/symmetric-tree/
// Related Topics 树
// 👍 60 👎 0
