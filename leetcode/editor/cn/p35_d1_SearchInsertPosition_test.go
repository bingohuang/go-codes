// github.com/bingohuang/go-codes
/**
题目: 35. 搜索插入位置
难度: 1
地址: https://leetcode-cn.com/problems/search-insert-position/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO35 struct {
	in1 []int
	in2 int
	out int
}

func Test35(t *testing.T) {
	// add test cases
	tc := map[string]IO35{
		"1": IO35{[]int{1, 3, 5, 6}, 5, 2},
		"2": IO35{[]int{1, 3, 5, 6}, 2, 1},
		"3": IO35{[]int{1, 3, 5, 6}, 7, 4},
		"4": IO35{[]int{1, 3, 5, 6}, 0, 0},
	}

	for k, v := range tc {
		// algo func
		out := searchInsert(v.in1, v.in2)

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
func searchInsert(nums []int, target int) int {
	index, begin, end := -1, 0, len(nums)
	for index == -1 {
		mid := (begin + end) / 2
		if target == nums[mid] {
			index = mid
		} else if target < nums[mid] {
			if mid == 0 || target > nums[mid-1] {
				index = mid
			}
			end = mid - 1
		} else {
			if mid == len(nums)-1 || target < nums[mid+1] {
				index = mid + 1
			}
			begin = mid + 1
		}
	}

	return index
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 你可以假设数组中无重复元素。
//
// 示例 1:
//
// 输入: [1,3,5,6], 5
//输出: 2
//
//
// 示例 2:
//
// 输入: [1,3,5,6], 2
//输出: 1
//
//
// 示例 3:
//
// 输入: [1,3,5,6], 7
//输出: 4
//
//
// 示例 4:
//
// 输入: [1,3,5,6], 0
//输出: 0
//
// Related Topics 数组 二分查找
