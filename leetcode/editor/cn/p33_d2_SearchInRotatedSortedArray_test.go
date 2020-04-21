// github.com/bingohuang/go-codes
/**
题目: 33. 搜索旋转排序数组
难度: 2
地址: https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO33 struct {
	in1 []int
	in2 int
	out int
}

func Test33(t *testing.T) {
	// add test cases
	tc := map[string]IO33{
		"1": IO33{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		"2": IO33{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		"3": IO33{[]int{4, 5, 6, 7, 0, 1, 2}, 7, 3},
		"4": IO33{[]int{4, 5, 6, 7, 0, 1, 2}, 4, 0},
		"5": IO33{[]int{4, 5, 6, 7, 0, 1, 2}, 2, 6},
		"6": IO33{[]int{5, 1, 3}, 3, 2},
	}

	for k, v := range tc {
		// algo func
		out := search(v.in1, v.in2)

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
func search(nums []int, target int) int {
	var binarySearch func(nums []int, target, low, high int) int
	binarySearch = func(nums []int, target, low, high int) int {
		if low > high {
			return -1
		}
		middle := low + (high-low)/2
		if nums[middle] == target {
			return middle
		}
		if nums[low] <= nums[middle] {
			if nums[low] <= target && target < nums[middle] {
				return binarySearch(nums, target, low, middle-1)
			}
			return binarySearch(nums, target, middle+1, high)
		} else {
			if nums[middle] < target && target <= nums[high] {
				return binarySearch(nums, target, middle+1, high)
			}
			return binarySearch(nums, target, low, middle-1)
		}
	}

	return binarySearch(nums, target, 0, len(nums)-1)
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//
// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
//
// 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
//
// 你可以假设数组中不存在重复的元素。
//
// 你的算法时间复杂度必须是 O(log n) 级别。
//
// 示例 1:
//
// 输入: nums = [4,5,6,7,0,1,2], target = 0
//输出: 4
//
//
// 示例 2:
//
// 输入: nums = [4,5,6,7,0,1,2], target = 3
//输出: -1
// Related Topics 数组 二分查找
