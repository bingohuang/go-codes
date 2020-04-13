// github.com/bingohuang/go-codes
/**
题目: 213. 打家劫舍 II
难度: 2
地址: https://leetcode-cn.com/problems/house-robber-ii/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO213 struct {
	in  []int
	out int
}

func Test213(t *testing.T) {
	// add test cases
	tc := map[string]IO213{
		"1": IO213{[]int{2, 3, 2}, 3},
		"2": IO213{[]int{1, 2, 3, 1}, 4},
	}

	for k, v := range tc {
		// algo func
		out := rob2(v.in)

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
// TODO: 提交请修改函数名为rob
func rob2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max2(rob2Range(nums, 0, n-2),
		rob2Range(nums, 1, n-1))
}

func rob2Range(nums []int, start, end int) int {
	dp_i, dp_i_1, dp_i_2 := 0, 0, 0
	for i := end; i >= start; i-- {
		dp_i = max2(dp_i_1, nums[i]+dp_i_2)
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i
}

func max2(x, y int) int {
	if x < y {
		return y
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋
//装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
//
// 示例 1:
//
// 输入: [2,3,2]
//输出: 3
//解释: 你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
//
//
// 示例 2:
//
// 输入: [1,2,3,1]
//输出: 4
//解释: 你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
//     偷窃到的最高金额 = 1 + 3 = 4 。
// Related Topics 动态规划
