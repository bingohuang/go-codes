// github.com/bingohuang/go-codes
/**
题目: 11. 盛最多水的容器
难度: 2
地址: https://leetcode-cn.com/problems/container-with-most-water/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO11 struct {
	in  []int
	out int
}

func Test11(t *testing.T) {
	// add test cases
	tc := map[string]IO11{
		"0": {[]int{}, 0},
		"1": {[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}

	for k, v := range tc {
		// algo func
		out := maxArea(v.in)

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
func maxArea(height []int) int {
	// 20200806
	// 1、双指针法
	// 时间复杂度：O(N)，双指针总计最多遍历整个数组一次。
	// 空间复杂度：O(1)，只需要额外的常数级别的空间。
	// 执行耗时:20 ms,击败了29.33% 的Go用户
	// 内存消耗:5.8 MB,击败了50.47% 的Go用户
	l, r := 0, len(height)-1
	ans := 0
	min := func(x, y int) int {
		if x <= y {
			return x
		}
		return y
	}
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		ans = max(ans, area)
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i,
//ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 说明：你不能倾斜容器，且 n 的值至少为 2。
//
//
//
//
//
// 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
//
//
// 示例：
//
// 输入：[1,8,6,2,5,4,8,3,7]
//输出：49
// Related Topics 数组 双指针
// 👍 1701 👎 0
