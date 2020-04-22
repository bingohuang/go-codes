// github.com/bingohuang/go-codes
/**
题目: 52. N皇后 II
难度: 3
地址: https://leetcode-cn.com/problems/n-queens-ii/
*/
package test

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// input and ouput
type IO52 struct {
	in  int
	out int
}

func Test52(t *testing.T) {
	// add test cases
	tc := map[string]IO52{
		"1": IO52{3, 0},
		"2": IO52{4, 2},
		"3": IO52{5, 10},
	}

	for k, v := range tc {
		// algo func
		out := totalNQueens(v.in)

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
func totalNQueens(n int) int {
	// 20200310:
	// 执行耗时:4 ms
	// 内存消耗:2.6 MB
	//res52 = [][]string{}
	//chessBoard := make([][]bool, n)
	//for i := 0; i < n; i++ {
	//	chessBoard[i] = make([]bool, n)
	//}
	//trackBack52(chessBoard, [][]byte{})
	//return len(res52)

	// 20200422：考虑做一版优化
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2 MB,击败了100.00% 的Go用户
	count := 0
	// 从第一行开始直到第 row 行的前一行为止，看那一行所放置的皇后是否在 col 列上，或者是不是在它的对角线上
	check := func(row, col int, columns []int) bool {
		for r := 0; r < row; r++ {
			if columns[r] == col || row-r == int(math.Abs(float64(columns[r]-col))) {
				return false
			}
		}
		return true
	}
	var backtracking func(n, row int, columns []int)
	backtracking = func(n, row int, columns []int) {
		// 是否在所有n行里都摆放好了皇后？
		if row == n {
			count++ // 找到了新的摆放方法
			return
		}
		// 尝试着将皇后放置在当前行中的每一列
		for col := 0; col < n; col++ {
			columns[row] = col

			// 检查是否合法，如果合法就继续到下一行
			if check(row, col, columns) {
				backtracking(n, row+1, columns)
			}
			// 如果不合法，就不要把皇后放在这列中（回溯）
			columns[row] = -1
		}
	}
	columns := make([]int, n)
	backtracking(n, 0, columns)

	return count
}

var res52 [][]string

func trackBack52(chessBoard [][]bool, track [][]byte) {
	if len(track) == len(chessBoard) {
		t := make([]string, len(track))
		for k, bs := range track {
			t[k] = string(bs)
		}
		res52 = append(res52, t)
	}

	for j := 0; j < len(chessBoard); j++ {
		if !valid52(chessBoard, len(track), j) {
			continue
		}
		bs := getLine52(len(chessBoard))
		bs[j] = 'Q'
		chessBoard[len(track)][j] = true
		track = append(track, bs)
		trackBack52(chessBoard, track)
		track = track[:len(track)-1]
		chessBoard[len(track)][j] = false
	}
}

func valid52(chessBoard [][]bool, row, cow int) bool {
	var i, j int
	for i = 0; i < row; i++ {
		if chessBoard[i][cow] == true {
			return false
		}
	}

	for j = 0; j < len(chessBoard); j++ {
		if chessBoard[row][j] == true {
			return false
		}
	}

	i, j = row, cow
	for i >= 0 && j >= 0 {
		if chessBoard[i][j] == true {
			return false
		}
		i--
		j--
	}

	i, j = row, cow
	for i >= 0 && j < len(chessBoard) {
		if chessBoard[i][j] == true {
			return false
		}
		j++
		i--
	}
	return true
}

func getLine52(n int) []byte {
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = '.'
	}
	return bs
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
//
//
// 上图为 8 皇后问题的一种解法。
//
// 给定一个整数 n，返回 n 皇后不同的解决方案的数量。
//
// 示例:
//
// 输入: 4
//输出: 2
//解释: 4 皇后问题存在如下两个不同的解法。
//[
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // 解法 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]
//
// Related Topics 回溯算法
