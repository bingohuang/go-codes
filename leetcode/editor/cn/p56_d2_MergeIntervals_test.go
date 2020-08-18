// github.com/bingohuang/go-codes
/**
题目: 56. 合并区间
难度: 2
地址: https://leetcode-cn.com/problems/merge-intervals/
*/
package test

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"testing"
)

// input and ouput
type IO56 struct {
	in  [][]int
	out [][]int
}

func Test56(t *testing.T) {
	// add test cases
	tc := map[string]IO56{
		"1": {[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		"2": {[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
	}

	for k, v := range tc {
		// algo func
		out := merge56(v.in)

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
func merge56(intervals [][]int) [][]int {
	// 20200423
	// 执行耗时:12 ms,击败了81.35% 的Go用户
	// 内存消耗:4.7 MB,击败了100.00% 的Go用户
	// 将所有的区间按照起始时间的先后顺序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 定义一个 previous 变量，初始化为 null
	var prev []int
	// 定义一个 result 变量，用来保存最终的区间结果
	var res [][]int
	// 从头开始遍历给定的所有区间
	for _, curr := range intervals {
		// 如果这是第一个区间，或者当前区间和前一个区间没有重叠，那么将当前区间加入到结果中
		if prev == nil || curr[0] > prev[1] {
			prev = curr
			res = append(res, prev)
		} else {
			// 否则，两个区间发生了重叠，更新前一个区间的结束时间
			prev[1] = int(math.Max(float64(prev[1]), float64(curr[1])))
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给出一个区间的集合，请合并所有重叠的区间。
//
// 示例 1:
//
// 输入: [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
//
// 示例 2:
//
// 输入: [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
// Related Topics 排序 数组
