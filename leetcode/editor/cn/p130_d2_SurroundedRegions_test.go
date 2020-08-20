// github.com/bingohuang/go-codes
/**
题目: 130. 被围绕的区域
难度: 2
地址: https://leetcode-cn.com/problems/surrounded-regions/
日期：2020-08-20 14:23:32
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO130 struct {
	in  [][]byte
	out [][]byte
}

func Test130(t *testing.T) {
	// add test cases
	tc := map[string]IO130{
		"0": {nil, nil},
		"1": {[][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		}, [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'X', 'X'},
		}},
	}

	for k, v := range tc {
		// algo func

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput: %v\n", v.in)
		solve(v.in)
		fmt.Printf("\toutput: %v\n", v.in)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(v.in, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, v.in)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func solve(board [][]byte) {
	// 2020-08-20 15:12 @bingohuang
	// 算法：1、深度优先搜索
	// 复杂度：O(MxN)/O(MxN)
	// 效率：执行耗时:36 ms,击败了9.94% 的Go用户
	//			内存消耗:6 MB,击败了38.39% 的Go用户
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	r, c := len(board), len(board[0])
	var dfs func([][]byte, int, int)
	dfs = func(board [][]byte, x, y int) {
		if x < 0 || x >= r || y < 0 || y >= c || board[x][y] != 'O' {
			return
		}
		board[x][y] = 'A'
		dfs(board, x+1, y)
		dfs(board, x-1, y)
		dfs(board, x, y+1)
		dfs(board, x, y-1)
	}
	// 左右两列搜索
	for i := 0; i < r; i++ {
		dfs(board, i, 0)
		dfs(board, i, c-1)
	}

	// 上下两行搜索
	for i := 0; i < c; i++ {
		dfs(board, 0, i)
		dfs(board, r-1, i)
	}

	// 对 A 标记，重新标记为O，对O标记为X
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
//
// 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
//
// 示例:
//
// X X X X
//X O O X
//X X O X
//X O X X
//
//
// 运行你的函数后，矩阵变为：
//
// X X X X
//X X X X
//X X X X
//X O X X
//
//
// 解释:
//
// 被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被
//填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
// Related Topics 深度优先搜索 广度优先搜索 并查集
// 👍 358 👎 0
