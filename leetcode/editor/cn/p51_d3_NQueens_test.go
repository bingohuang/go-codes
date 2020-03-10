// github.com/bingohuang/go-codes
/**
题目: 51. N皇后
难度: 3
地址: https://leetcode-cn.com/problems/n-queens/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO51 struct {
	in  int
	out [][]string
}

func Test51(t *testing.T) {
	// add test cases
	tc := map[string]IO51{
		"1": IO51{2, [][]string{}},
		"2": IO51{3, [][]string{}},
		"3": IO51{4, [][]string{
			{".Q..",
				"...Q",
				"Q...",
				"..Q."},
			{"..Q.",
				"Q...",
				"...Q",
				".Q.."},
		}},
	}

	for k, v := range tc {
		// algo func
		out := solveNQueens(v.in)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput: %v\n", v.in)
		fmt.Printf("\toutput: %v\n", out)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

func solveNQueens1(n int) [][]string {
	res = [][]string{}
	chessBoard := make([][]bool, n)
	for i := 0; i < n; i++ {
		chessBoard[i] = make([]bool, n)
	}
	trackBack1(chessBoard, [][]byte{})
	return res
}

func trackBack1(chessBoard [][]bool, track [][]byte) {
	if len(track) == len(chessBoard) {
		t := make([]string, len(track))
		for k, bs := range track {
			t[k] = string(bs)
		}
		res = append(res, t)
	}

	for j := 0; j < len(chessBoard); j++ {
		if !valid(chessBoard, len(track), j) {
			continue
		}
		bs := getLine(len(chessBoard))
		bs[j] = 'Q'
		chessBoard[len(track)][j] = true
		track = append(track, bs)
		trackBack1(chessBoard, track)
		track = track[:len(track)-1]
		chessBoard[len(track)][j] = false
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
var res [][]string

func solveNQueens(n int) [][]string {
	res = [][]string{}
	chessBoard := make([][]bool, n)
	for i := 0; i < n; i++ {
		chessBoard[i] = make([]bool, n)
	}
	trackBack(chessBoard, [][]byte{})
	return res
}

func trackBack(chessBoard [][]bool, track [][]byte) {
	if len(track) == len(chessBoard) {
		t := make([]string, len(track))
		for k, bs := range track {
			t[k] = string(bs)
		}
		res = append(res, t)
	}

	for j := 0; j < len(chessBoard); j++ {
		if !valid(chessBoard, len(track), j) {
			continue
		}
		bs := getLine(len(chessBoard))
		bs[j] = 'Q'
		chessBoard[len(track)][j] = true
		track = append(track, bs)
		trackBack(chessBoard, track)
		track = track[:len(track)-1]
		chessBoard[len(track)][j] = false
	}
}

func valid(chessBoard [][]bool, row, cow int) bool {
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

func getLine(n int) []byte {
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
// 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
//
// 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
// 示例:
//
// 输入: 4
//输出: [
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // 解法 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]
//解释: 4 皇后问题存在两个不同的解法。
//
// Related Topics 回溯算法
