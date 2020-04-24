// github.com/bingohuang/go-codes
/**
题目: 778. 水位上升的泳池中游泳
难度: 3
地址: https://leetcode-cn.com/problems/swim-in-rising-water/
*/
package test

import (
	"container/heap"
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO778 struct {
	in  [][]int
	out int
}

func Test778(t *testing.T) {
	// add test cases
	tc := map[string]IO778{
		"1": {[][]int{{0, 2}, {1, 3}}, 3},
	}

	for k, v := range tc {
		// algo func
		out := swimInWater(v.in)

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
func swimInWater(grid [][]int) int {
	// 20200424
	// 执行耗时:12 ms,击败了86.11% 的Go用户
	// 内存消耗:5.7 MB,击败了100.00% 的Go用户
	const maxN = 50
	n := len(grid)
	dr := []int{1, -1, 0, 0}
	dc := []int{0, 0, 1, -1}
	visited := [maxN][maxN]bool{}
	result := 0
	pq := &posHeap{}
	heap.Push(pq, pos{height: grid[0][0], r: 0, c: 0})
	visited[0][0] = true

	for pq.Len() > 0 {
		info := heap.Pop(pq).(pos) // 游到当前最低的平台上
		if grid[info.r][info.c] > result {
			result = grid[info.r][info.c]
		}
		if info.r == n-1 && info.c == n-1 { // 终点
			break
		}
		for i := 0; i < len(dr); i++ {
			r, c := info.r+dr[i], info.c+dc[i]
			if r >= 0 && r < n && c >= 0 && c < n && !visited[r][c] {
				heap.Push(pq, pos{height: grid[r][c], r: r, c: c})
				visited[r][c] = true
			}
		}
	}
	return result
}

// 平台结构体，方便自定义heap实现
type pos struct {
	height, r, c int // 高度和横纵坐标
}

type posHeap []pos

func (h posHeap) Len() int            { return len(h) }
func (h posHeap) Less(i, j int) bool  { return h[i].height < h[j].height }
func (h posHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *posHeap) Push(x interface{}) { *h = append(*h, x.(pos)) }
func (h *posHeap) Pop() interface{} {
	pos := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return pos
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//在一个 N x N 的坐标方格 grid 中，每一个方格的值 grid[i][j] 表示在位置 (i,j) 的平台高度。
//
// 现在开始下雨了。当时间为 t 时，此时雨水导致水池中任意位置的水位为 t 。你可以从一个平台游向四周相邻的任意一个平台，但是前提是此时水位必须同时淹没这两
//个平台。假定你可以瞬间移动无限距离，也就是默认在方格内部游动是不耗时的。当然，在你游泳的时候你必须待在坐标方格里面。
//
// 你从坐标方格的左上平台 (0，0) 出发。最少耗时多久你才能到达坐标方格的右下平台 (N-1, N-1)？
//
// 示例 1:
//
//
//输入: [[0,2],[1,3]]
//输出: 3
//解释:
//时间为0时，你位于坐标方格的位置为 (0, 0)。
//此时你不能游向任意方向，因为四个相邻方向平台的高度都大于当前时间为 0 时的水位。
//
//等时间到达 3 时，你才可以游向平台 (1, 1). 因为此时的水位是 3，坐标方格中的平台没有比水位 3 更高的，所以你可以游向坐标方格中的任意位置
//
//
// 示例2:
//
//
//输入: [[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6
//]]
//输入: 16
//解释:
// 0  1  2  3  4
//24 23 22 21  5
//12 13 14 15 16
//11 17 18 19 20
//10  9  8  7  6
//
//最终的路线用加粗进行了标记。
//我们必须等到时间为 16，此时才能保证平台 (0, 0) 和 (4, 4) 是连通的
//
//
// 提示:
//
//
// 2 <= N <= 50.
// grid[i][j] 位于区间 [0, ..., N*N - 1] 内。
//
// Related Topics 堆 深度优先搜索 并查集 二分查找
