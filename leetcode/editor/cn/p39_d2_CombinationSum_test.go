// github.com/bingohuang/go-codes
/**
题目: 39. 组合总和
难度: 2
地址: https://leetcode-cn.com/problems/combination-sum/
*/
package test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// input and ouput
type IO39 struct {
	in1 []int
	in2 int
	out [][]int
}

func Test39(t *testing.T) {
	// add test cases
	tc := map[string]IO39{
		"0": {[]int{2, 3}, 7, [][]int{{2, 2, 3}}},
		"1": {[]int{2, 3, 6, 7}, 7, [][]int{{7}, {2, 2, 3}}},
		"2": {[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		"3": {[]int{8, 7, 4, 3}, 11, [][]int{{3, 4, 4}, {3, 8}, {4, 7}}},
	}

	for k, v := range tc {
		// algo func
		out := combinationSum(v.in1, v.in2)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput1: %v\n", v.in1)
		fmt.Printf("\tinput1: %v\n", v.in2)
		fmt.Printf("\toutput: %v\n", out)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) [][]int {
	// 20200422
	res := make([][]int, 0)
	sol := make([]int, 0)
	// 必须排序
	sort.Ints(candidates)
	// 不去重
	// 执行耗时:4 ms,击败了81.68% 的Go用户
	// 内存消耗:2.7 MB,击败了100.00% 的Go用户

	// 去重
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.7 MB,击败了100.00% 的Go用户
	nums := removeDuplicates(candidates)

	backtracking(nums, target, 0, sol, &res)
	return res
}
func backtracking(candidates []int, target int, start int, sol []int, res *[][]int) {
	if target == 0 {
		tmp := make([]int, len(sol))
		copy(tmp, sol)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if target-candidates[i] < 0 {
			break
		}
		sol = append(sol, candidates[i])
		backtracking(candidates, target-candidates[i], i, sol, res)
		sol = sol[:len(sol)-1]
	}
}

func removeDuplicates(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return nums[0 : j+1]
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的数字可以无限制重复被选取。
//
// 说明：
//
//
// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。
//
//
// 示例 1:
//
// 输入: candidates = [2,3,6,7], target = 7,
//所求解集为:
//[
//  [7],
//  [2,2,3]
//]
//
//
// 示例 2:
//
// 输入: candidates = [2,3,5], target = 8,
//所求解集为:
//[
//  [2,2,2,2],
//  [2,3,3],
//  [3,5]
//]
// Related Topics 数组 回溯算法
