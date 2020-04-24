// github.com/bingohuang/go-codes
/**
题目: 490. 迷宫
难度: 2
地址: https://leetcode-cn.com/problems/the-maze/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO490 struct {
	in1 [][]int
	in2 []int
	in3 []int
	out bool
}

func Test490(t *testing.T) {
	// add test cases
	tc := map[string]IO490{
		"1": {[][]int{
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0},
			{1, 1, 0, 1, 1},
			{0, 0, 0, 0, 0},
		}, []int{0, 4}, []int{4, 4}, true},
		"2": {[][]int{
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0},
			{1, 1, 0, 1, 1},
			{0, 0, 0, 0, 0},
		}, []int{0, 4}, []int{3, 2}, false},
	}

	for k, v := range tc {
		// algo func
		out := hasPath(v.in1, v.in2, v.in3)

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
func hasPath(maze [][]int, start []int, destination []int) bool {
	// 20200424
	// DFS: O(MN)/O(MN)
	// 执行耗时:36 ms,击败了100.00% 的Go用户
	// 内存消耗:6.6 MB,击败了100.00% 的Go用户
	//visited := make([][]bool, len(maze))
	//for i := range visited {
	//	visited[i] = make([]bool, len(maze[0]))
	//}
	//var dfs func(maze [][]int, start []int, destination []int, visited [][]bool) bool
	//dfs = func(maze [][]int, start []int, destination []int, visited [][]bool) bool {
	//	if visited[start[0]][start[1]] {
	//		return false
	//	}
	//	if start[0] == destination[0] && start[1] == destination[1] {
	//		return true
	//	}
	//	visited[start[0]][start[1]] = true
	//	u, d, l, r := start[0] - 1, start[0] + 1, start[1] - 1, start[1] + 1
	//	// 一直往上探索
	//	for u >=0 && maze[u][start[1]] == 0 {
	//		u --
	//	}
	//	if dfs(maze, []int{u+1, start[1]}, destination, visited) {
	//		return true
	//	}
	//	// 一直往下探索
	//	for d < len(maze) && maze[d][start[1]] == 0 {
	//		d ++
	//	}
	//	if dfs(maze, []int{d-1, start[1]}, destination, visited) {
	//		return true
	//	}
	//	// 一直往左探索
	//	for l >=0 && maze[start[0]][l] == 0 {
	//		l --
	//	}
	//	if dfs(maze, []int{start[0], l + 1}, destination, visited) {
	//		return true
	//	}
	//	// 一直往右探索
	//	for r < len(maze[0]) && maze[start[0]][r] == 0 {
	//		r ++
	//	}
	//	if dfs(maze, []int{start[0], r - 1}, destination, visited) {
	//		return true
	//	}
	//	return false
	//}
	//return dfs(maze, start, destination, visited)

	// BFS: O(MN)/O(MN)
	// 执行耗时:40 ms,击败了100.00% 的Go用户
	// 内存消耗:6.4 MB,击败了100.00% 的Go用户
	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[0]))
	}
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	queue := [][]int{start}
	visited[start[0]][start[1]] = true
	for len(queue) > 0 {
		// 从队列里取一个元素
		s := queue[0]
		queue = queue[1:]
		if s[0] == destination[0] && s[1] == destination[1] {
			return true
		}
		for _, dir := range dirs {
			x := s[0] + dir[0]
			y := s[1] + dir[1]
			for x >= 0 && x < len(maze) && y >= 0 && y < len(maze[0]) && maze[x][y] == 0 {
				x += dir[0]
				y += dir[1]
			}
			if !visited[x-dir[0]][y-dir[1]] {
				queue = append(queue, []int{x - dir[0], y - dir[1]})
				visited[x-dir[0]][y-dir[1]] = true
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//由空地和墙组成的迷宫中有一个球。球可以向上下左右四个方向滚动，但在遇到墙壁前不会停止滚动。当球停下时，可以选择下一个方向。
//
// 给定球的起始位置，目的地和迷宫，判断球能否在目的地停下。
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
//输出: true
//
//解析: 一个可能的路径是 : 左 -> 下 -> 左 -> 下 -> 右 -> 下 -> 右。
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
//输出: false
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
