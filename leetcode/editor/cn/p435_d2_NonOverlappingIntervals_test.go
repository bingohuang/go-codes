// github.com/bingohuang/go-codes
/**
题目: 435. 无重叠区间
难度: 2
地址: https://leetcode-cn.com/problems/non-overlapping-intervals/
*/
package test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// input and ouput
type IO435 struct {
	in  [][]int
	out int
}

func Test435(t *testing.T) {
	// add test cases
	tc := map[string]IO435{
		"1": {[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		"2": {[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		"3": {[][]int{{1, 2}, {2, 3}}, 0},
	}

	for k, v := range tc {
		// algo func
		out := eraseOverlapIntervals(v.in)

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
func eraseOverlapIntervals(intervals [][]int) int {
	// 20200423
	// 解题思路一： 暴力法

	// 解题思路二：贪婪法-按照开始时间排序

	// 解题思路三：贪婪法-按照结束时间排序
	// 执行耗时:8 ms,击败了90.30% 的Go用户
	// 内存消耗:4.1 MB,击败了100.00% 的Go用户
	// 先算出最多有多少个区间相互之间是没有重叠的，则最少需要将 n-m 个区间删除就行
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	end := intervals[0][1]
	count := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			end = intervals[i][1]
			count++
		}
	}
	return len(intervals) - count
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
//
// 注意:
//
//
// 可以认为区间的终点总是大于它的起点。
// 区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
//
//
// 示例 1:
//
//
//输入: [ [1,2], [2,3], [3,4], [1,3] ]
//
//输出: 1
//
//解释: 移除 [1,3] 后，剩下的区间没有重叠。
//
//
// 示例 2:
//
//
//输入: [ [1,2], [1,2], [1,2] ]
//
//输出: 2
//
//解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
//
//
// 示例 3:
//
//
//输入: [ [1,2], [2,3] ]
//
//输出: 0
//
//解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
//
// Related Topics 贪心算法
