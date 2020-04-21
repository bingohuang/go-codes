// github.com/bingohuang/go-codes
/**
题目: 34. 在排序数组中查找元素的第一个和最后一个位置
难度: 2
地址: https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO34 struct {
	in1 []int
	in2 int
	out []int
}

func Test34(t *testing.T) {
	// add test cases
	tc := map[string]IO34{
		"1": IO34{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		"2": IO34{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		"3": IO34{[]int{5, 7, 7, 8, 8, 10}, 5, []int{0, 0}},
		"4": IO34{[]int{5, 7, 7, 8, 8, 10}, 10, []int{5, 5}},
	}

	for k, v := range tc {
		// algo func
		out := searchRange(v.in1, v.in2)

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
func searchRange(nums []int, target int) []int {

	var searchLowerBound func(nums []int, target, low, high int) int
	var searchUpperBound func(nums []int, target, low, high int) int

	searchLowerBound = func(nums []int, target, low, high int) int {
		if low > high {
			return -1
		}
		middle := low + (high-low)/2

		if nums[middle] == target && (middle == 0 || nums[middle-1] < target) {
			return middle
		}
		if target <= nums[middle] {
			return searchLowerBound(nums, target, low, middle-1)
		} else {
			return searchLowerBound(nums, target, middle+1, high)
		}
	}

	searchUpperBound = func(nums []int, target, low, high int) int {
		if low > high {
			return -1
		}
		middle := low + (high-low)/2

		if nums[middle] == target && (middle == len(nums)-1 || nums[middle+1] > target) {
			return middle
		}
		if target < nums[middle] {
			return searchUpperBound(nums, target, low, middle-1)
		} else {
			return searchUpperBound(nums, target, middle+1, high)
		}
	}
	lower := searchLowerBound(nums, target, 0, len(nums)-1)
	upper := searchUpperBound(nums, target, 0, len(nums)-1)

	return []int{lower, upper}
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
// 你的算法时间复杂度必须是 O(log n) 级别。
//
// 如果数组中不存在目标值，返回 [-1, -1]。
//
// 示例 1:
//
// 输入: nums = [5,7,7,8,8,10], target = 8
//输出: [3,4]
//
// 示例 2:
//
// 输入: nums = [5,7,7,8,8,10], target = 6
//输出: [-1,-1]
// Related Topics 数组 二分查找
