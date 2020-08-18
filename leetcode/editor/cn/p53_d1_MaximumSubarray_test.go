// github.com/bingohuang/go-codes
/**
题目: 53. 最大子序和
难度: 1
地址: https://leetcode-cn.com/problems/maximum-subarray/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO53 struct {
	in  []int
	out int
}

func Test53(t *testing.T) {
	// add test cases
	tc := map[string]IO53{
		//"0": {[]int{}, 0},
		"1": {[]int{1}, 1},
		"2": {[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	for k, v := range tc {
		// algo func
		out := maxSubArray(v.in)

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
func maxSubArray(nums []int) int {
	// 20200818
	// 1、暴力法：找到所有的连续子数组，判定最大值
	// Time Limit Exceeded
	/*if len(nums) == 0 {
		return 0
	}
	max := math.MinInt64
	n := len(nums)
	for i := 0; i < n;i++ {
		for j:= 1; j <= n;j ++ {
			sum := 0
			for k := i; k < i + j && k < n; k ++ {
				sum += nums[k]
			}
			if sum > max {
				max = sum
			}
		}
	}
	return max*/
	// 2、动态规划
	// f(i)=max{f(i-1)+ai,ai}
	// O(N)/O(1)
	// 执行耗时:8 ms,击败了58.98% 的Go用户
	// 内存消耗:3.3 MB,击败了72.41% 的Go用户
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 示例:
//
// 输入: [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//
//
// 进阶:
//
// 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
// Related Topics 数组 分治算法 动态规划
// 👍 2308 👎 0
