// github.com/bingohuang/go-codes
/**
题目: 253. 会议室 II
难度: 2
地址: https://leetcode-cn.com/problems/meeting-rooms-ii/
*/
package test

import (
	"container/heap"
	"reflect"
	"sort"
	"testing"
)

// input and ouput
type IO253 struct {
	in  [][]int
	out int
}

func Test253(t *testing.T) {
	// add test cases
	tc := map[string]IO253{
		"1": {[][]int{{0, 30}, {5, 10}, {15, 20}}, 2},
		//"2": {[][]int{{7, 10}, {2, 4}}, 1},
		//"3": {[][]int{{7, 10}, {10, 12}}, 1},
	}

	for k, v := range tc {
		// algo func
		out := minMeetingRooms(v.in)
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func minMeetingRooms(intervals [][]int) int {
	// 20200423
	// 执行耗时:8 ms,击败了94.17% 的Go用户
	// 内存消耗:5.3 MB,击败了100.00% 的Go用户

	if len(intervals) == 0 {
		return 0
	}

	// 将输入的一系列会议按照会议的起始时间排序。
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})

	// 用一个最小堆来维护目前开辟的所有会议室，最小堆里的会议室按照会议的结束时间排序。
	h := &IntHeap{intervals[0]}
	// 让第一个会议在第一个会议室里举行。
	heap.Init(h)

	for i := 1; i < len(intervals); i++ {
		// 从第二个会议开始，对于每个会议，我们都从最小堆里取出一个会议室，那么这个会议室里的会议一定是最早结束的。
		firstInterval := heap.Pop(h).([]int)
		if intervals[i][0] >= firstInterval[1] {
			// 若当前要开的会议可以等会议室被腾出才开始，那么就可以重复利用这个会议室。
			firstInterval[1] = intervals[i][1]
		} else {
			// 否则，开一个新的会议室。
			heap.Push(h, intervals[i])
		}
		// 把旧的会议室也放入到最小堆里。
		heap.Push(h, firstInterval)
	}
	// 最小堆里的会议室个数就是要求的答案，即最少的会议个数。
	return h.Len()

}

type IntHeap [][]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][1] < h[j][1] } // 最小堆里的会议室按照会议的结束时间排序。
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个会议时间安排的数组，每个会议时间都会包括开始和结束的时间 [[s1,e1],[s2,e2],...] (si < ei)，为避免会议冲突，同时要考虑
//充分利用会议室资源，请你计算至少需要多少间会议室，才能满足这些会议安排。
//
// 示例 1:
//
// 输入: [[0, 30],[5, 10],[15, 20]]
//输出: 2
//
// 示例 2:
//
// 输入: [[7,10],[2,4]]
//输出: 1
// Related Topics 堆 贪心算法 排序
