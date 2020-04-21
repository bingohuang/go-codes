// github.com/bingohuang/go-codes
/**
题目: [1]两数之和
难度: 1
地址: https://leetcode-cn.com/problems/two-sum/
*/
package test

import (
	"reflect"
	"testing"
)

// input and ouput
type IO1 struct {
	in1 []int
	in2 int
	out []int
}

func Test1(t *testing.T) {
	// add test cases
	tc := map[string]IO1{
		"1": IO1{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		"2": IO1{[]int{2, 7, 2, 15}, 4, []int{0, 2}},
		"3": IO1{[]int{1, 2, 2, 6}, 4, []int{1, 2}},
	}

	for k, v := range tc {
		// algo func
		out := twoSum1(v.in1, v.in2)
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

// 算法1：暴力破解法
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 算法2：map法
func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//
// 示例:
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
// Related Topics 数组 哈希表
