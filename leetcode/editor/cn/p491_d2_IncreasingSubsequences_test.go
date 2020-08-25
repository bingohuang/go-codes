// github.com/bingohuang/go-codes
/**
题目: 491. 递增子序列
难度: 2
地址: https://leetcode-cn.com/problems/increasing-subsequences/
日期：2020-08-25 19:21:01
*/
package test

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// input and ouput
type IO491 struct {
	in  []int
	out [][]int
}

func Test491(t *testing.T) {

	// add test cases
	tc := map[string]IO491{
		//"0": {[]int{0}, [][]int{{}}},
		"1": {[]int{4, 6, 7, 7}, [][]int{
			{4, 6}, {4, 7}, {4, 6, 7}, {4, 6, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}, {4, 7, 7},
		}},
	}

	for k, v := range tc {
		// algo func
		out := findSubsequences(v.in)

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
func findSubsequences(nums []int) [][]int {
	// 2020-08-25 19:48 @bingohuang
	// 算法：1、递归枚举+剪枝
	// 复杂度：O(2^n)*n/O(n)
	// 效率：执行耗时:28 ms,击败了88.64% 的Go用户
	//			内存消耗:12.5 MB,击败了76.92% 的Go用户
	var (
		ans [][]int
		tmp []int
		dfs func([]int, int, int)
	)
	dfs = func(nums []int, cur, last int) {
		if cur == len(nums) {
			if len(tmp) >= 2 {
				t := make([]int, len(tmp))
				copy(t, tmp)
				ans = append(ans, t)
			}
			return
		}
		if nums[cur] >= last {
			tmp = append(tmp, nums[cur])
			dfs(nums, cur+1, nums[cur])
			tmp = tmp[:len(tmp)-1]
		}
		if nums[cur] != last {
			dfs(nums, cur+1, last)
		}
	}
	dfs(nums, 0, math.MinInt32)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
//
// 示例:
//
//
//输入: [4, 6, 7, 7]
//输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7
//]]
//
// 说明:
//
//
// 给定数组的长度不会超过15。
// 数组中的整数范围是 [-100,100]。
// 给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
//
// Related Topics 深度优先搜索
// 👍 168 👎 0
