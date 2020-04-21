// github.com/bingohuang/go-codes
/**
题目: [252]会议室
难度: 1
地址: https://leetcode-cn.com/problems/meeting-rooms
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
//Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei),
// determine if a person could attend all meetings.
//
//Example 1:
//
//Input: [[0,30],[5,10],[15,20]]
//Output: false
//Example 2:
//
//Input: [[7,10],[2,4]]
//Output: true
//NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.
