// github.com/bingohuang/go-codes
/**
题目: 22. 括号生成
难度: 2
地址: https://leetcode-cn.com/problems/generate-parentheses/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO22 struct {
	in  int
	out []string
}

func Test22(t *testing.T) {
	// add test cases
	tc := map[string]IO22{
		"0": {0, nil},
		"1": {1, []string{"()"}},
		"2": {3, []string{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()"}},
	}

	for k, v := range tc {
		// algo func
		out := generateParenthesis(v.in)

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
func generateParenthesis(n int) []string {
	// 20200817
	// 1、递归找规律
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.8 MB,击败了32.43% 的Go用户
	//
	if n <= 0 {
		return nil
	}
	var res []string
	var getParenthesis func(str string, left, right int)
	getParenthesis = func(str string, left, right int) {
		// 如果左右括号都用完，则添加到结果中，并退出
		if left == 0 && right == 0 {
			res = append(res, str)
			return
		}
		if left == right {
			// 剩余左右括号数相等，下一个只能用左括号
			getParenthesis(str+"(", left-1, right)
		} else if left < right {
			// 剩余左括号小于右括号，下一个可以用左括号也可以用右括号
			if left > 0 {
				getParenthesis(str+"(", left-1, right)
			}
			getParenthesis(str+")", left, right-1)
		}
	}
	getParenthesis("", n, n)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例：
//
// 输入：n = 3
//输出：[
//       "((()))",
//       "(()())",
//       "(())()",
//       "()(())",
//       "()()()"
//     ]
//
// Related Topics 字符串 回溯算法
// 👍 1241 👎 0
