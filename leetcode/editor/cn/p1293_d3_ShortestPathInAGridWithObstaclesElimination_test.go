// github.com/bingohuang/go-codes
/**
题目: 1293. 网格中的最短路径
难度: 3
地址: https://leetcode-cn.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/
*/
package test

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO1293 struct {
	in1 [][]int
	in2 int
	out int
}

func Test1293(t *testing.T) {
	// add test cases
	tc := map[string]IO1293{
		"1": {[][]int{
			{0, 0, 0},
			{1, 1, 0},
			{0, 0, 0},
			{0, 1, 1},
			{0, 0, 0},
		}, 1, 6},
		//"2": {[][]int{
		//	{0, 1, 1},
		//	{1, 1, 1},
		//	{1, 0, 0},
		//}, 1, -1},
		//"3": {[][]int{ // 构造容易超时的用例
		//	{0,1,0,0,1,1,1,0,1,1,0,0,1,1,1,0,1},
		//	{0,1,0,0,0,0,0,1,0,0,1,0,0,1,0,1,1},
		//	{1,1,0,1,1,1,1,1,0,0,0,0,1,0,0,1,1},
		//	{0,0,0,1,1,1,1,1,0,0,1,0,1,1,1,0,1},
		//	{1,0,1,0,0,1,0,0,1,0,1,0,0,0,1,1,1},
		//	{0,0,1,0,1,0,1,0,1,0,0,1,1,1,1,1,0},
		//	{1,0,1,1,0,0,1,0,0,1,1,1,1,1,1,0,1},
		//	{0,1,1,1,1,0,0,1,0,1,1,0,0,1,1,1,0},
		//	{1,1,0,0,1,1,1,0,1,0,1,1,1,0,1,1,1},
		//	{0,0,0,1,0,1,1,1,1,1,1,1,0,0,0,0,0},
		//	{0,1,0,1,0,0,1,1,0,1,0,1,0,0,0,0,1},
		//	{1,0,1,0,1,0,1,0,0,0,1,1,1,0,0,0,1},
		//	{1,1,0,0,0,0,1,1,1,1,0,0,0,1,1,1,0},
		//	{0,0,1,1,1,1,1,1,0,1,0,0,1,1,0,1,1},
		//	{0,0,1,0,1,1,0,0,0,1,0,1,0,0,0,0,0},
		//	{0,1,1,0,1,1,0,1,1,1,0,1,0,1,1,1,1},
		//	{0,0,0,1,1,1,1,1,0,0,1,1,1,0,1,0,1},
		//	{0,1,1,0,1,0,0,0,1,1,1,0,1,1,1,0,0},
		//	{0,0,0,1,1,1,0,1,1,1,0,1,1,0,1,0,0},
		//	{0,0,1,0,0,1,0,1,1,0,0,1,0,1,1,0,0},
		//	{0,1,1,0,0,1,0,0,0,1,1,1,1,0,0,1,1},
		//	{1,0,1,0,1,0,0,0,1,0,1,0,0,0,0,0,1},
		//	{1,1,1,0,1,1,0,1,0,0,1,1,0,0,1,1,1},
		//	{1,0,0,0,1,1,0,1,1,1,0,1,1,1,1,1,0},
		//}, 38, 39},
	}

	for k, v := range tc {
		// algo func
		out := shortestPath(v.in1, v.in2)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput1: %v\n", v.in1)
		fmt.Printf("\tinput2: %v\n", v.in2)
		fmt.Printf("\toutput: %v\n", out)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func shortestPath(grid [][]int, k int) int {
	// 20200424
	// DFS
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.5 MB,击败了100.00% 的Go用户
	//const maxSize = 40
	//m, n := len(grid), len(grid[0])
	//if k >= m + n-3 { // 这里不能少了=号，否则会Time Limit Exceeded
	//	return m + n -2
	//}
	//passed := 0
	//visited := [maxSize][maxSize]bool{}
	//var dfs func(r, c, k int) int
	//dfs = func(r, c, k int) int {
	//	if r < 0 || r >= m || c < 0 || c >= n ||
	//		visited[r][c] || grid[r][c] == 1 && k == 0 {
	//		return math.MaxInt32
	//	}
	//	if r == m-1 && c == n-1 {
	//		return passed
	//	}
	//	passed ++
	//	visited[r][c] = true
	//	if grid[r][c] == 1 {
	//		k--
	//	}
	//	left := dfs(r, c-1, k)
	//	right := dfs(r, c+1, k)
	//	up := dfs(r-1, c, k)
	//	down := dfs(r+1, c, k)
	//	passed --
	//	visited[r][c] = false
	//
	//	return min(min(left, right), min(up, down))
	//}
	//res := dfs(0,0,k)
	//if res == math.MaxInt32 {
	//	return -1
	//}
	//return res

	// 20200424
	// BFS
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.7 MB,击败了100.00% 的Go用户
	const maxSize = 40
	const maxK = maxSize + maxSize - 3
	m, n := len(grid), len(grid[0])
	if k >= m+n-3 {
		return m + n - 2
	}
	visited := [maxSize][maxSize][maxK + 1]bool{}
	queue := list.New()
	queue.PushBack([]int{0, 0, k})
	visited[0][0][k] = true
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	passed := 0
	for queue.Len() > 0 {
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ {
			info := queue.Remove(queue.Front()).([]int)
			r, c, k := info[0], info[1], info[2]
			if r == m-1 && c == n-1 {
				return passed
			}
			if grid[r][c] == 1 {
				k--
			}
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr < 0 || nr >= m || nc < 0 || nc >= n ||
					visited[nr][nc][k] || grid[nr][nc] == 1 && k == 0 {
					continue
				}
				queue.PushBack([]int{nr, nc, k})
				visited[nr][nc][k] = true
			}
		}
		passed++
	}
	return -1
}

//
//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}
//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给你一个 m * n 的网格，其中每个单元格不是 0（空）就是 1（障碍物）。每一步，您都可以在空白单元格中上、下、左、右移动。
//
// 如果您 最多 可以消除 k 个障碍物，请找出从左上角 (0, 0) 到右下角 (m-1, n-1) 的最短路径，并返回通过该路径所需的步数。如果找不到这样
//的路径，则返回 -1。
//
//
//
// 示例 1：
//
// 输入：
//grid =
//[[0,0,0],
// [1,1,0],
// [0,0,0],
// [0,1,1],
// [0,0,0]],
//k = 1
//输出：6
//解释：
//不消除任何障碍的最短路径是 10。
//消除位置 (3,2) 处的障碍后，最短路径是 6 。该路径是 (0,0) -> (0,1) -> (0,2) -> (1,2) -> (2,2) -> (3
//,2) -> (4,2).
//
//
//
//
// 示例 2：
//
// 输入：
//grid =
//[[0,1,1],
// [1,1,1],
// [1,0,0]],
//k = 1
//输出：-1
//解释：
//我们至少需要消除两个障碍才能找到这样的路径。
//
//
//
//
// 提示：
//
//
// grid.length == m
// grid[0].length == n
// 1 <= m, n <= 40
// 1 <= k <= m*n
// grid[i][j] == 0 or 1
// grid[0][0] == grid[m-1][n-1] == 0
//
// Related Topics 广度优先搜索
