// github.com/bingohuang/go-codes
/**
题目: 300. 最长上升子序列
难度: 2
地址: https://leetcode-cn.com/problems/longest-increasing-subsequence/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO300 struct {
	in  []int
	out int
}

func Test300(t *testing.T) {
	// add test cases
	tc := map[string]IO300{
		//"1": IO300{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		"2": IO300{[]int{1, 3, 2, 3, 1, 4}, 4},
		//"3": IO300{[]int{4, 10, 4, 3, 8, 9}, 3},
	}

	for k, v := range tc {
		// algo func
		out := lengthOfLIS(v.in)

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
func lengthOfLIS(nums []int) int {
	// 20200422-栈实现
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.4 MB,击败了75.00% 的Go用户
	if len(nums) == 0 {
		return 0
	}
	stack := make([]int, 0, len(nums))
	stack = append(stack, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] > stack[len(stack)-1] {
			stack = append(stack, nums[i])
		} else {
			for j := 0; j < len(stack); j++ {
				if stack[j] >= nums[i] {
					stack[j] = nums[i]
					break
				}
			}
		}
	}
	return len(stack)

	// 20200422-优化
	// 执行耗时:8 ms,击败了68.83% 的Go用户
	// 内存消耗:2.4 MB,击败了50.00% 的Go用户
	//if len(nums) == 0 {
	//	return 0
	//}
	//dp := make([]int, len(nums))
	//dp[0] = 1
	//lis := 1
	//for i := 1; i < len(nums); i++ {
	//	dp[i] = 1
	//	for j := 0; j < i; j++ {
	//		if nums[i] > nums[j] && dp[i] < dp[j]+1 {
	//			dp[i] = dp[j] + 1
	//		}
	//	}
	//	if lis < dp[i] {
	//		lis = dp[i]
	//	}
	//}
	//return lis

	// 20200403 O(N^2)
	//dp := make([]int, len(nums))
	//for i := range dp {
	//	dp[i] = 1
	//}
	//for i := 0; i < len(nums); i++ {
	//	for j := 0; j < i; j++ {
	//		if nums[i] > nums[j] {
	//			if dp[i] < dp[j]+1 {
	//				dp[i] = dp[j] + 1
	//			}
	//		}
	//	}
	//}
	//res := 0
	//for _, v := range dp {
	//	if v > res {
	//		res = v
	//	}
	//}
	//return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个无序的整数数组，找到其中最长上升子序列的长度。
//
// 示例:
//
// 输入: [10,9,2,5,3,7,101,18]
//输出: 4
//解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
//
// 说明:
//
//
// 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
// 你算法的时间复杂度应该为 O(n2) 。
//
//
// 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
// Related Topics 二分查找 动态规划
