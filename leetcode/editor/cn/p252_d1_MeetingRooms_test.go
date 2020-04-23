// github.com/bingohuang/go-codes
/**
题目: 252. 会议室
难度: 1
地址: https://leetcode-cn.com/problems/meeting-rooms/
*/
package test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// input and ouput
type IO252 struct {
	in  [][]int
	out bool
}

func Test252(t *testing.T) {
	// add test cases
	tc := map[string]IO252{
		"1": {[][]int{{0, 30}, {5, 10}, {15, 20}}, false},
		"2": {[][]int{{7, 10}, {2, 4}}, true},
		"3": {[][]int{{7, 10}, {10, 12}}, true},
	}

	for k, v := range tc {
		// algo func
		out := canAttendMeetings(v.in)
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func canAttendMeetings(intervals [][]int) bool {
	// 20200422
	// 执行耗时:16 ms,击败了31.37% 的Go用户
	// 内存消耗:6.2 MB,击败了100.00% 的Go用户
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	fmt.Printf("intervals=%v\n", intervals)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true

}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个会议时间安排的数组，每个会议时间都会包括开始和结束的时间 [[s1,e1],[s2,e2],...] (si < ei)，请你判断一个人是否能够参加
//这里面的全部会议。
//
// 示例 1:
//
// 输入: [[0,30],[5,10],[15,20]]
//输出: false
//
//
// 示例 2:
//
// 输入: [[7,10],[2,4]]
//输出: true
//
// Related Topics 排序
