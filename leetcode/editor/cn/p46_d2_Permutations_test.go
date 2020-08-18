// github.com/bingohuang/go-codes
/**
题目: 46. 全排列
难度: 2
地址: https://leetcode-cn.com/problems/permutations/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO46 struct {
	in  []int
	out [][]int
}

func Test46(t *testing.T) {
	// add test cases
	tc := map[string]IO46{
		//"0": {[]int{}, nil},
		//"1": {[]int{1}, [][]int{{1}}},
		"2": {[]int{1, 2}, [][]int{{1, 2}, {2, 1}}},
		"3": {[]int{1, 2, 3}, [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}},
	}

	for k, v := range tc {
		// algo func
		out := permute(v.in)

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
func permute(nums []int) [][]int {
	// 20200818
	// https://leetcode-cn.com/problems/permutations/solution/quan-pai-lie-by-leetcode-solution-2/
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.6 MB,击败了89.64% 的Go用户
	var res [][]int
	n := len(nums)
	var backtrack func(first int)
	backtrack = func(first int) {
		if first == n {
			// 注意这里必须用tmp+copy，否则结果全重复
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}
		for i := first; i < n; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}

	}
	backtrack(0)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
// 示例:
//
// 输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//]
// Related Topics 回溯算法
// 👍 845 👎 0
