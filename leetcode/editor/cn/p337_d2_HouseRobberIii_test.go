// github.com/bingohuang/go-codes
/**
题目: 337. 打家劫舍 III
难度: 2
地址: https://leetcode-cn.com/problems/house-robber-iii/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO337 struct {
	in  *TreeNode
	out int
}

func Test337(t *testing.T) {
	// add test cases
	treeNode_in1 := &TreeNode{
		3,
		&TreeNode{
			2,
			nil,
			&TreeNode{
				3,
				nil,
				nil,
			},
		},
		&TreeNode{
			3,
			nil,
			&TreeNode{
				1,
				nil,
				nil,
			},
		},
	}
	treeNode_in2 := &TreeNode{
		3,
		&TreeNode{
			4,
			&TreeNode{
				1,
				nil,
				nil,
			},
			&TreeNode{
				3,
				nil,
				nil,
			},
		},
		&TreeNode{
			5,
			nil,
			&TreeNode{
				1,
				nil,
				nil,
			},
		},
	}
	tc := map[string]IO337{
		"1": IO337{treeNode_in1, 7},
		"2": IO337{treeNode_in2, 9},
	}

	for k, v := range tc {
		// algo func
		out := rob3(v.in)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput: %v\n", v.in)
		fmt.Printf("\toutput: %v\n", out)
		fmt.Printf("\texcept: %v\n", v.out)
		fmt.Printf("\tmemoMap: %v\n", memoMap)

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
// TODO: 提交代码修改函数名
func rob3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 利用备忘录消除子问题
	if v, ok := memoMap[root]; ok {
		return v
	}
	// 抢，然后去下下家
	do_it := root.Val
	if root.Left != nil {
		do_it += rob3(root.Left.Left) + rob3(root.Left.Right)
	}
	if root.Right != nil {
		do_it += rob3(root.Right.Left) + rob3(root.Right.Right)
	}
	// 不抢，然后去下家
	not_do := rob3(root.Left) + rob3(root.Right)

	res := max(do_it, not_do)
	memoMap[root] = res

	return res
}

// TODO: 提交代码打开注释
//var memoMap = make(map[*TreeNode]int)
//
//func max(x, y int) int {
//	if x < y {
//		return y
//	}
//	return x
//}
//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“
//房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
//
// 计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。
//
// 示例 1:
//
// 输入: [3,2,3,null,3,null,1]
//
//     3
//    / \
//   2   3
//    \   \
//     3   1
//
//输出: 7
//解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
//
// 示例 2:
//
// 输入: [3,4,5,1,3,null,1]
//
//     3
//    / \
//   4   5
//  / \   \
// 1   3   1
//
//输出: 9
//解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.
//
// Related Topics 树 深度优先搜索
