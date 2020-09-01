// github.com/bingohuang/go-codes
/**
题目: 486. 预测赢家
难度: 2
地址: https://leetcode-cn.com/problems/predict-the-winner/
日期：2020-09-01 17:03:08
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO486 struct {
	in  []int
	out bool
}

func Test486(t *testing.T) {
	// add test cases
	tc := map[string]IO486{
		"0": {[]int{1}, true},
		"1": {[]int{1, 5, 2}, false},
		"2": {[]int{1, 5, 233, 7}, true},
	}

	for k, v := range tc {
		// algo func
		out := PredictTheWinner(v.in)

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
func PredictTheWinner(nums []int) bool {
	// 2020-09-01 17:05 @bingohuang
	// 算法：1、动态规划：
	// d[i][j]表示当前数组剩下的部分为下标i到下标j时
	// 当前玩家与另一位玩家的分数之差最大值，注意当前玩家不一定是先手
	// if i > j dp[i][j] = 0
	// if i == j dp[i][j] = nums[i]
	// if i < j dp[i][j]=max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
	// check dp[0][len(nums)-1] >= 0
	// 复杂度1：执行耗时:0 ms,击败了100.00% 的Go用户
	//			内存消耗:2.1 MB,击败了13.04% 的Go用户
	// 效率：O(N^2)/O(N^2)
	l := len(nums)
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	/*dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
		dp[i][i] = nums[i]
	}
	for i := l - 2; i >=0; i-- {
		for j := i + 1; j < l; j ++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][l-1] >= 0*/
	// 优化存储空间
	// 复杂度：O(N^2)/O(N)
	// 效率：执行耗时:0 ms,击败了100.00% 的Go用户
	//			内存消耗:2 MB,击败了52.17% 的Go用户
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = nums[i]
	}
	for i := l - 2; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			dp[j] = max(nums[i]-dp[j], nums[j]-dp[j-1])
		}
	}
	return dp[l-1] >= 0
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个表示分数的非负整数数组。 玩家 1 从数组任意一端拿取一个分数，随后玩家 2 继续从剩余数组任意一端拿取分数，然后玩家 1 拿，…… 。每次一个玩家
//只能拿取一个分数，分数被拿取之后不再可取。直到没有剩余分数可取时游戏结束。最终获得分数总和最多的玩家获胜。
//
// 给定一个表示分数的数组，预测玩家1是否会成为赢家。你可以假设每个玩家的玩法都会使他的分数最大化。
//
//
//
// 示例 1：
//
// 输入：[1, 5, 2]
//输出：False
//解释：一开始，玩家1可以从1和2中进行选择。
//如果他选择 2（或者 1 ），那么玩家 2 可以从 1（或者 2 ）和 5 中进行选择。如果玩家 2 选择了 5 ，那么玩家 1 则只剩下 1（或者 2 ）
//可选。
//所以，玩家 1 的最终分数为 1 + 2 = 3，而玩家 2 为 5 。
//因此，玩家 1 永远不会成为赢家，返回 False 。
//
//
// 示例 2：
//
// 输入：[1, 5, 233, 7]
//输出：True
//解释：玩家 1 一开始选择 1 。然后玩家 2 必须从 5 和 7 中进行选择。无论玩家 2 选择了哪个，玩家 1 都可以选择 233 。
//     最终，玩家 1（234 分）比玩家 2（12 分）获得更多的分数，所以返回 True，表示玩家 1 可以成为赢家。
//
//
//
//
// 提示：
//
//
// 1 <= 给定的数组长度 <= 20.
// 数组里所有分数都为非负数且不会大于 10000000 。
// 如果最终两个玩家的分数相等，那么玩家 1 仍为赢家。
//
// Related Topics 极小化极大 动态规划
// 👍 272 👎 0
