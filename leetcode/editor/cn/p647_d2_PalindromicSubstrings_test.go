// github.com/bingohuang/go-codes
/**
题目: 647. 回文子串
难度: 2
地址: https://leetcode-cn.com/problems/palindromic-substrings/
日期: 2020-08-19 19:32:27
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO647 struct {
	in  string
	out int
}

func Test647(t *testing.T) {
	// add test cases
	tc := map[string]IO647{
		"0": IO647{"", 0},
		"1": IO647{"abc", 3},
		"2": IO647{"aaa", 6},
	}

	for k, v := range tc {
		// algo func
		out := countSubstrings(v.in)

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
func countSubstrings(s string) int {
	// 20200819
	// 1、数学法
	// https://leetcode-cn.com/problems/palindromic-substrings/solution/hui-wen-zi-chuan-by-leetcode-solution/
	// O(n^2)/O(1)
	// Runtime:0 ms, faster than 100.00% of Go online submissions.
	// Memory Usage:1.9 MB, less than 100.00% of Go online submissions.
	n := len(s)
	ans := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
//
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
//
//
//
// 示例 1：
//
// 输入："abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
//
//
// 示例 2：
//
// 输入："aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
//
//
//
// 提示：
//
//
// 输入的字符串长度不会超过 1000 。
//
// Related Topics 字符串 动态规划
// 👍 342 👎 0
