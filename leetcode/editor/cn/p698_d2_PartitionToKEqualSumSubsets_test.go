// github.com/bingohuang/go-codes
/**
题目: 698. 划分为k个相等的子集
难度: 2
地址: https://leetcode-cn.com/problems/partition-to-k-equal-sum-subsets/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO698 struct {
	in1 []int
	in2 int
	out bool
}

func Test698(t *testing.T) {
	// add test cases
	tc := map[string]IO698{
		"1": {[]int{4, 3, 2, 3, 5, 2, 1}, 4, true},
	}

	for k, v := range tc {
		// algo func
		out := canPartitionKSubsets(v.in1, v.in2)

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
func canPartitionKSubsets(nums []int, k int) bool {
	// 20200424
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.1 MB,击败了50.00% 的Go用户
	if len(nums) < k {
		return false
	}
	sum, max := 0, 0 // 注意到输入限制为非负，且和不会溢出
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	target := sum / k // 尝试把所有元素放到k个组中，每组元素和为target
	if sum%k != 0 || max > target {
		return false
	}
	used := make([]bool, len(nums))
	return backTracking698(k, 0, 0, target, nums, used)
}

/*
尝试把所有元素放到k个组中，每组元素和为target
currSubSum 记录当前组累加的结果，start指明从nums的哪个位置开始
*/
func backTracking698(k, currSubSum, start, target int, nums []int, used []bool) bool {
	if k == 0 { // 说明所有数字都放入了对应组
		return true
	}
	if currSubSum == target { // 已经构建了若干组
		// 构建下一组
		return backTracking698(k-1, 0, 0, target, nums, used)
	}
	for i := start; i < len(nums); i++ {
		if !used[i] && currSubSum+nums[i] <= target {
			used[i] = true // 当前值放入组
			if backTracking698(k, currSubSum+nums[i], i+1, target, nums, used) {
				return true
			}
			used[i] = false // 说明将当前值放入一组不能得到结果，回溯
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个整数数组 nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
//
// 示例 1：
//
//
//输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
//输出： True
//说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
//
//
//
// 注意:
//
//
// 1 <= k <= len(nums) <= 16
// 0 < nums[i] < 10000
//
// Related Topics 递归 动态规划
