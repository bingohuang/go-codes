/*
	题目大意：
		有N个任务需要处理，第i个任务开始时间si，结束时间ti。
		现在你需要开一些线程来处理问题，这些线程只能串行处理任务，即分配给一个线程的所有任务，任意两个任务的执行时间区间不能相交。
		额外的，一个任务刚结束后可以立马开始下一个任务。
		现在你要计算最少需开多少个线程，才能处理完所有任务。

	PS：题目未给出每个任务的开始结束时间的大小，假定范围在int32内。

	样例1：
		输入：[[1,7],[4,10]]
		输出：2
	样例2：
		输入：[[2,7],[7,10]]
		输出：1
*/

package main

import (
	"fmt"
	"sort"
)

type sortType [][]int

func (s sortType) Len() int {
	return len(s)
}

func (s sortType) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortType) Less(i, j int) bool {
	if s[i][0] != s[j][0] {
		return s[i][0] < s[j][0]
	}
	return s[i][1] < s[j][1] // 结束时间和开始时间相同时，结束时间放前面
}

func minThreadsInPool(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	var segPos [][]int
	for i := 0; i < n; i++ {
		segPos = append(segPos, []int{intervals[i][0], 1})
		segPos = append(segPos, []int{intervals[i][1], -1})
	}
	sort.Sort(sortType(segPos))
	ans := 0
	now := 0
	for i := 0; i < len(segPos); i++ {
		now += segPos[i][1]
		if now > ans {
			ans = now
		}
	}
	return ans
}

func main() {
	intervals := [][]int{{7, 10}, {2, 7}}
	ans := minThreadsInPool(intervals)
	fmt.Println(ans)
}
