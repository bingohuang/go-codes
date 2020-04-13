// github.com/bingohuang/go-codes
/**
题目: 198. 打家劫舍
难度: 1
地址: https://leetcode-cn.com/problems/house-robber/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO198 struct {
	in  []int
	out int
}

func Test198(t *testing.T) {
	// add test cases
	tc := map[string]IO198{
		"1": IO198{[]int{1, 2, 3, 1}, 4},
		"2": IO198{[]int{2, 7, 9, 3, 1}, 12},
	}

	for k, v := range tc {
		// algo func
		out := rob1(v.in)

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
func rob1(nums []int) int {
	// 1
	//return dp1(nums, 0)

	// 2
	//memo = make([]int, len(nums))
	//// 这个初始化必要
	//for i := range memo {
	//	memo[i] = -1
	//}
	//return dp2(nums, 0)

	// 3
	//n := len(nums)
	//dp := make([]int, n+2)
	//for i := n-1;i>=0;i--{
	//	dp[i] = max(dp[i+1], nums[i] + dp[i+2])
	//}
	//return dp[0]

	// 4
	n := len(nums)
	dp_i, dp_i_1, dp_i_2 := 0, 0, 0
	for i := n - 1; i >= 0; i-- {
		dp_i = max1(dp_i_1, nums[i]+dp_i_2)
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i
}

func max1(x, y int) int {
	if x < y {
		return y
	}
	return x
}

var memo []int

func dp2(nums []int, start int) int {
	if start >= len(nums) {
		return 0
	}
	if memo[start] != -1 {
		return memo[start]
	}
	res1 := dp2(nums, start+1)
	res2 := nums[start] + dp2(nums, start+2)
	if res2 > res1 {
		res1 = res2
	}
	memo[start] = res1
	return res1
}

// Time Limit Exceeded
func dp1(nums []int, start int) int {
	if start >= len(nums) {
		return 0
	}
	rob := nums[start] + dp1(nums, start+2)
	norob := dp1(nums, start+1)
	if rob > norob {
		return rob
	}
	return norob
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上
//被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
//
// 示例 1:
//
// 输入: [1,2,3,1]
//输出: 4
//解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 2:
//
// 输入: [2,7,9,3,1]
//输出: 12
//解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
//
// Related Topics 动态规划
