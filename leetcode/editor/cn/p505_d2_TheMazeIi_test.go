// github.com/bingohuang/go-codes
/**
题目: 505. 迷宫 II
难度: 2
地址: https://leetcode-cn.com/problems/the-maze-ii/
*/
package test

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// input and ouput
type IO505 struct {
	in1 [][]int
	in2 []int
	in3 []int
	out int
}

func Test505(t *testing.T) {
	// add test cases
	tc := map[string]IO505{
		"1": {[][]int{
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0},
			{1, 1, 0, 1, 1},
			{0, 0, 0, 0, 0},
		}, []int{0, 4}, []int{4, 4}, 12},
		"2": {[][]int{
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0},
			{1, 1, 0, 1, 1},
			{0, 0, 0, 0, 0},
		}, []int{0, 4}, []int{3, 2}, -1},
	}

	for k, v := range tc {
		// algo func
		out := shortestDistance(v.in1, v.in2, v.in3)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput1: %v\n", v.in1)
		fmt.Printf("\tinput2: %v\n", v.in2)
		fmt.Printf("\tinput3: %v\n", v.in3)
		fmt.Printf("\toutput: %v\n", out)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func shortestDistance(maze [][]int, start []int, destination []int) int {
	// 20200424
	// DFS:O(MN*max(M,N))/O(MN)
	// 执行耗时:540 ms,击败了8.33% 的Go用户
	//// 内存消耗:7.5 MB,击败了100.00% 的Go用户
	//distance := make([][]int, len(maze))
	//for i := range distance {
	//	distance[i] = make([]int, len(maze[0]))
	//	for j := range distance[i] {
	//		distance[i][j] = math.MaxInt32
	//	}
	//}
	//distance[start[0]][start[1]] = 0
	//var dfs func(maze [][]int, start []int, destination []int)
	//dfs = func(maze [][]int, start []int, destination []int) {
	//	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	//	for _, dir := range dirs {
	//		x := start[0] + dir[0]
	//		y := start[1] + dir[1]
	//		count := 0
	//		for x >= 0 && x < len(maze) && y >= 0 && y < len(maze[0]) && maze[x][y] == 0 {
	//			x += dir[0]
	//			y += dir[1]
	//			count ++
	//		}
	//		if distance[start[0]][start[1]] + count < distance[x - dir[0]][y - dir[1]] {
	//			distance[x - dir[0]][y - dir[1]] = distance[start[0]][start[1]] + count
	//			dfs(maze, []int{x - dir[0], y - dir[1]},destination)
	//		}
	//	}
	//}
	//dfs(maze, start,destination)
	//if distance[destination[0]][destination[1]] == math.MaxInt32 {
	//	return -1
	//}
	//return distance[destination[0]][destination[1]]

	// 20200424
	// BFS:O(MN*max(M,N))/O(MN)
	// 执行耗时:44 ms,击败了91.67% 的Go用户
	// 内存消耗:6.5 MB,击败了100.00% 的Go用户
	distance := make([][]int, len(maze))
	for i := range distance {
		distance[i] = make([]int, len(maze[0]))
		for j := range distance[i] {
			distance[i][j] = math.MaxInt32
		}
	}
	distance[start[0]][start[1]] = 0
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	queue := [][]int{start}
	for len(queue) != 0 {
		s := queue[0]
		queue = queue[1:]
		for _, dir := range dirs {
			x := s[0] + dir[0]
			y := s[1] + dir[1]
			count := 0
			for x >= 0 && x < len(maze) && y >= 0 && y < len(maze[0]) && maze[x][y] == 0 {
				x += dir[0]
				y += dir[1]
				count++
			}
			if distance[s[0]][s[1]]+count < distance[x-dir[0]][y-dir[1]] {
				distance[x-dir[0]][y-dir[1]] = distance[s[0]][s[1]] + count
				queue = append(queue, []int{x - dir[0], y - dir[1]})
			}
		}
	}

	if distance[destination[0]][destination[1]] == math.MaxInt32 {
		return -1
	}
	return distance[destination[0]][destination[1]]
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//由空地和墙组成的迷宫中有一个球。球可以向上下左右四个方向滚动，但在遇到墙壁前不会停止滚动。当球停下时，可以选择下一个方向。
//
// 给定球的起始位置，目的地和迷宫，找出让球停在目的地的最短距离。距离的定义是球从起始位置（不包括）到目的地（包括）经过的空地个数。如果球无法停在目的地，返回
// -1。
//
// 迷宫由一个0和1的二维数组表示。 1表示墙壁，0表示空地。你可以假定迷宫的边缘都是墙壁。起始位置和目的地的坐标通过行号和列号给出。
//
//
//
// 示例 1:
//
// 输入 1: 迷宫由以下二维数组表示
//
//0 0 1 0 0
//0 0 0 0 0
//0 0 0 1 0
//1 1 0 1 1
//0 0 0 0 0
//
//输入 2: 起始位置坐标 (rowStart, colStart) = (0, 4)
//输入 3: 目的地坐标 (rowDest, colDest) = (4, 4)
//
//输出: 12
//
//解析: 一条最短路径 : left -> down -> left -> down -> right -> down -> right。
//             总距离为 1 + 1 + 3 + 1 + 2 + 2 + 2 = 12。
//
//
//
// 示例 2:
//
// 输入 1: 迷宫由以下二维数组表示
//
//0 0 1 0 0
//0 0 0 0 0
//0 0 0 1 0
//1 1 0 1 1
//0 0 0 0 0
//
//输入 2: 起始位置坐标 (rowStart, colStart) = (0, 4)
//输入 3: 目的地坐标 (rowDest, colDest) = (3, 2)
//
//输出: -1
//
//解析: 没有能够使球停在目的地的路径。
//
//
//
//
//
// 注意:
//
//
// 迷宫中只有一个球和一个目的地。
// 球和目的地都在空地上，且初始时它们不在同一位置。
// 给定的迷宫不包括边界 (如图中的红色矩形), 但你可以假设迷宫的边缘都是墙壁。
// 迷宫至少包括2块空地，行数和列数均不超过100。
//
// Related Topics 深度优先搜索 广度优先搜索
