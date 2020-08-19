// github.com/bingohuang/go-codes
/**
题目: 79. 单词搜索
难度: 2
地址: https://leetcode-cn.com/problems/word-search/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO79 struct {
	in1 [][]byte
	in2 string
	out bool
}

func Test79(t *testing.T) {
	// add test cases
	// board =
	//[
	//  ['A','B','C','E'],
	//  ['S','F','C','S'],
	//  ['A','D','E','E']
	//]
	//
	//给定 word = "ABCCED", 返回 true
	//给定 word = "SEE", 返回 true
	//给定 word = "ABCB", 返回 false
	tc := map[string]IO79{
		//"0": {[][]byte{}, "",false},
		"1": {[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCCED", true},
		//"2": {[][]byte{
		//	{'A', 'B', 'C', 'E'},
		//	{'S', 'F', 'C', 'S'},
		//	{'A', 'D', 'E', 'E'},
		//}, "SEE", true},
		"3": {[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCB", false},
	}

	for k, v := range tc {
		// algo func
		out := exist(v.in1, v.in2)

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
func exist(board [][]byte, word string) bool {
	// 20200819
	// 1、回溯+DFS
	// https://leetcode-cn.com/problems/word-search/solution/hui-su-suan-fa-lei-si-ba-huang-hou-wen-ti-shuang-b/
	// O(4^n)~=O(2^n)
	// 执行耗时:4 ms,击败了98.57% 的Go用户
	// 内存消耗:3.5 MB,击败了61.86% 的Go用户
	var backtrace func(board [][]byte, word string, i, j, k int) bool
	backtrace = func(board [][]byte, word string, i, j, k int) bool {
		// 说明找到了，成功退出
		if k == len(word) {
			return true
		}
		// 超出边界，失败退出
		if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
			return false
		}
		// 字母不对，失败退出
		if board[i][j] != word[k] {
			return false
		}
		tmp := board[i][j]
		// 置空，以防重复使用
		board[i][j] = byte(' ')
		// 右
		if backtrace(board, word, i+1, j, k+1) {
			return true
		}
		// 下
		if backtrace(board, word, i, j-1, k+1) {
			return true
		}
		// 左
		if backtrace(board, word, i-1, j, k+1) {
			return true
		}
		// 上
		if backtrace(board, word, i, j+1, k+1) {
			return true
		}
		// 再赋值回去
		board[i][j] = tmp
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtrace(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个二维网格和一个单词，找出该单词是否存在于网格中。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
// 示例:
//
// board =
//[
//  ['A','B','C','E'],
//  ['S','F','C','S'],
//  ['A','D','E','E']
//]
//
//给定 word = "ABCCED", 返回 true
//给定 word = "SEE", 返回 true
//给定 word = "ABCB", 返回 false
//
//
//
// 提示：
//
//
// board 和 word 中只包含大写和小写英文字母。
// 1 <= board.length <= 200
// 1 <= board[i].length <= 200
// 1 <= word.length <= 10^3
//
// Related Topics 数组 回溯算法
// 👍 526 👎 0
