// github.com/bingohuang/go-codes
/**
题目: 78. 子集
难度: 2
地址: https://leetcode-cn.com/problems/subsets/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO78 struct {
	in  []int
	out [][]int
}

func Test78(t *testing.T) {
	// add test cases
	tc := map[string]IO78{
		"1": IO78{[]int{}, [][]int{
			{},
		}},
		"2": IO78{[]int{1}, [][]int{
			{},
			{1},
		}},
		"3": IO78{[]int{1, 2}, [][]int{
			{},
			{1},
			{1, 2},
			{2},
		}},
		"4": IO78{[]int{1, 2, 3}, [][]int{
			{},
			{1},
			{1, 2},
			{1, 2, 3},
			{1, 3},
			{2},
			{2, 3},
			{3},
		}},
	}

	for k, v := range tc {
		// algo func
		out := subsets(v.in)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput: %v\n", v.in)
		fmt.Printf("\toutput: %v\n", out)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

// 算法1：回溯1
func subsets1(nums []int) [][]int {
	result := make([][]int, 0)
	item := make([]int, 0)

	result = append(result, item)
	generate1(0, nums, &item, &result)
	return result
}

func generate1(i int, nums []int, item *[]int, result *[][]int) {
	if i >= len(nums) {
		return
	}
	*item = append(*item, nums[i])
	temp := make([]int, len(*item))
	for i, v := range *item {
		temp[i] = v
	}
	*result = append(*result, temp)
	generate1(i+1, nums, item, result)
	*item = (*item)[:len(*item)-1]
	generate1(i+1, nums, item, result)
	return
}

//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	item := make([]int, 0)

	result = append(result, item)
	generate(0, nums, &item, &result)
	return result
}

func generate(i int, nums []int, item *[]int, result *[][]int) {
	if i >= len(nums) {
		return
	}
	*item = append(*item, nums[i])
	temp := make([]int, len(*item))
	for i, v := range *item {
		temp[i] = v
	}
	*result = append(*result, temp)
	generate(i+1, nums, item, result)
	*item = (*item)[:len(*item)-1]
	generate(i+1, nums, item, result)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
// 说明：解集不能包含重复的子集。
//
// 示例:
//
// 输入: nums = [1,2,3]
//输出:
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]
// Related Topics 位运算 数组 回溯算法
