// github.com/bingohuang/go-codes
/**
题目: 55. 跳跃游戏
难度: 2
地址: https://leetcode-cn.com/problems/jump-game/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO55 struct {
	in  []int
	out bool
}

func Test55(t *testing.T) {
	// add test cases
	tc := map[string]IO55{
		"0": {[]int{}, true},
		"1": {[]int{2, 3, 1, 1, 4}, true},
		"2": {[]int{3, 2, 1, 0, 4}, false},
	}

	for k, v := range tc {
		// algo func
		out := canJump(v.in)

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
func canJump(nums []int) bool {
	// 20200818
	// 1、贪心法
	// 执行耗时:4 ms,击败了100.00% 的Go用户
	// 内存消耗:4.2 MB,击败了100.00% 的Go用户
	n := len(nums)
	rightmost := 0
	for i := 0; i < n; i++ {
		if i <= rightmost {
			if rightmost < i+nums[i] {
				rightmost = i + nums[i]
			}
			if rightmost >= n-1 {
				return true
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个非负整数数组，你最初位于数组的第一个位置。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 判断你是否能够到达最后一个位置。
//
// 示例 1:
//
// 输入: [2,3,1,1,4]
//输出: true
//解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
//
//
// 示例 2:
//
// 输入: [3,2,1,0,4]
//输出: false
//解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
//
// Related Topics 贪心算法 数组
// 👍 776 👎 0
