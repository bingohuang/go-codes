// github.com/bingohuang/go-codes
/**
题目: 473. 火柴拼正方形
难度: 2
地址: https://leetcode-cn.com/problems/matchsticks-to-square/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO473 struct {
	in  []int
	out bool
}

func Test473(t *testing.T) {
	// add test cases
	tc := map[string]IO473{
		"1": {[]int{1, 1, 2, 2, 2}, true},
		"2": {[]int{3, 3, 3, 3, 4}, false},
	}

	for k, v := range tc {
		// algo func
		out := makesquare(v.in)

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
func makesquare(nums []int) bool {
	// 20200424
	// 执行耗时:8 ms,击败了86.49% 的Go用户
	// 内存消耗:2.1 MB,击败了100.00% 的Go用户
	k := 4
	if len(nums) < 4 {
		return false
	}
	sum, max := 0, 0 // 注意到输入限制为非负，且和不会溢出
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	target := sum / k // 尝试把所有元素放到k个组中，每组元素和为target
	if sum%k != 0 || max > target {
		return false
	}
	used := make([]bool, len(nums))
	return backTracking473(k, 0, 0, target, nums, used)
}

/*
尝试把所有元素放到k个组中，每组元素和为target
currSubSum 记录当前组累加的结果，start指明从nums的哪个位置开始
*/
func backTracking473(k, currSubSum, start, target int, nums []int, used []bool) bool {
	if k == 0 { // 说明所有数字都放入了对应组
		return true
	}
	if currSubSum == target { // 已经构建了若干组
		// 构建下一组
		return backTracking473(k-1, 0, 0, target, nums, used)
	}
	for i := start; i < len(nums); i++ {
		if !used[i] && currSubSum+nums[i] <= target {
			used[i] = true // 当前值放入组
			if backTracking473(k, currSubSum+nums[i], i+1, target, nums, used) {
				return true
			}
			used[i] = false // 说明将当前值放入一组不能得到结果，回溯
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//还记得童话《卖火柴的小女孩》吗？现在，你知道小女孩有多少根火柴，请找出一种能使用所有火柴拼成一个正方形的方法。不能折断火柴，可以把火柴连接起来，并且每根火柴
//都要用到。
//
// 输入为小女孩拥有火柴的数目，每根火柴用其长度表示。输出即为是否能用所有的火柴拼成正方形。
//
// 示例 1:
//
//
//输入: [1,1,2,2,2]
//输出: true
//
//解释: 能拼成一个边长为2的正方形，每边两根火柴。
//
//
// 示例 2:
//
//
//输入: [3,3,3,3,4]
//输出: false
//
//解释: 不能用所有火柴拼成一个正方形。
//
//
// 注意:
//
//
// 给定的火柴长度和在 0 到 10^9之间。
// 火柴数组的长度不超过15。
//
// Related Topics 深度优先搜索
