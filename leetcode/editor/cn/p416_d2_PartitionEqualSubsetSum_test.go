// github.com/bingohuang/go-codes
/**
题目: 416. 分割等和子集
难度: 2
地址: https://leetcode-cn.com/problems/partition-equal-subset-sum/
*/
package test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// input and ouput
type IO416 struct {
	in  []int
	out bool
}

func Test416(t *testing.T) {
	// add test cases
	tc := map[string]IO416{
		"1": {[]int{1, 5, 11, 5}, true},
		"2": {[]int{1, 2, 3, 5}, false},
	}

	for k, v := range tc {
		// algo func
		out := canPartition(v.in)

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
func canPartition(nums []int) bool {
	// 20200424
	// Time Limit Exceeded
	//k := 2
	//if len(nums) < k {
	//	return false
	//}
	//sum, max := 0, 0 // 注意到输入限制为非负，且和不会溢出
	//for _, v := range nums {
	//	sum += v
	//	if v > max {
	//		max = v
	//	}
	//}
	//target := sum / k // 尝试把所有元素放到k个组中，每组元素和为target
	//if sum%k != 0 || max > target {
	//	return false
	//}
	//used := make([]bool, len(nums))
	//return backTracking416(k, 0, 0, target, nums, used)

	// 20200424
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.4 MB,击败了100.00% 的Go用户
	const groups = 2
	if len(nums) < groups {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	target := sum / groups
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	if sum%groups != 0 || nums[0] > target {
		return false
	}
	used := make([]bool, len(nums))
	return backTracking416(groups, 0, 0, target, nums, used)
}

/*
尝试把所有元素放到k个组中，每组元素和为target
currSubSum 记录当前组累加的结果，start指明从nums的哪个位置开始
*/
func backTracking416(k, currSubSum, start, target int, nums []int, used []bool) bool {
	if k == 0 { // 说明所有数字都放入了对应组
		return true
	}
	if currSubSum == target { // 已经构建了若干组
		// 构建下一组
		return backTracking416(k-1, 0, 0, target, nums, used)
	}
	for i := start; i < len(nums); i++ {
		if !used[i] && currSubSum+nums[i] <= target {
			used[i] = true // 当前值放入组
			if backTracking416(k, currSubSum+nums[i], i+1, target, nums, used) {
				return true
			}
			used[i] = false // 说明将当前值放入一组不能得到结果，回溯
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
// 注意:
//
//
// 每个数组中的元素不会超过 100
// 数组的大小不会超过 200
//
//
// 示例 1:
//
// 输入: [1, 5, 11, 5]
//
//输出: true
//
//解释: 数组可以分割成 [1, 5, 5] 和 [11].
//
//
//
//
// 示例 2:
//
// 输入: [1, 2, 3, 5]
//
//输出: false
//
//解释: 数组不能分割成两个元素和相等的子集.
//
//
//
// Related Topics 动态规划
