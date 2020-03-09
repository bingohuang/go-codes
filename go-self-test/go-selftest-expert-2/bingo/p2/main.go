package main

import (
	"fmt"
	"sort"
)

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
	样例3：
		输入：[[1,3],[2,4],[2,5]
		输出：3
*/
func main() {
	case1 := [][]int{{1, 7}, {4, 10}}
	fmt.Println("case1:", case1, "\n\tright:", 2, "\n\twrong:", minThead(case1))
	case2 := [][]int{{2, 7}, {7, 10}}
	fmt.Println("case2:", case2, "\n\tright:", 1, "\n\twrong:", minThead(case2))
	case3 := [][]int{{1, 3}, {2, 4}, {2, 5}}
	fmt.Println("case3:", case3, "\n\tright:", 3, "\n\twrong:", minThead(case3))
}

/*
思路1：
	将任务按照起始时间和结束时间展开，再对整个数组排序
	排序方式：首先按照启示时间排序，如果起始时间相同，则结束时间放前
*/
func minThead(tasks [][]int) int {
	n := len(tasks)
	if n == 0 {
		return 0
	}
	var times [][]int
	for i := 0; i < len(tasks); i++ {
		times = append(times, []int{tasks[i][0], 1})
		times = append(times, []int{tasks[i][1], -1})
	}
	fmt.Println("times:", times)
	sort.Sort(sortType(times))
	fmt.Println("times:", times)

	var now, ans int
	for i := 0; i < len(times); i++ {
		now += times[i][1]
		if now > ans {
			ans = now
		}
	}

	return ans
}

type sortType [][]int

func (s sortType) Len() int {
	return len(s)
}

func (s sortType) Less(i, j int) bool {
	if s[i][0] != s[j][0] {
		return s[i][0] < s[j][0]
	}
	return s[i][1] < s[j][1]
}

func (s sortType) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
