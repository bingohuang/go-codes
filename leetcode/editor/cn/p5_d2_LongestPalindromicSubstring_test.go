// github.com/bingohuang/go-codes
/**
题目: 5. 最长回文子串
难度: 2
地址: https://leetcode-cn.com/problems/longest-palindromic-substring/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO5 struct {
	in  string
	out string
}

func Test5(t *testing.T) {
	// add test cases
	tc := map[string]IO5{
		"0": IO5{"", ""},
		"1": IO5{"babad", "bab"},
		"2": IO5{"cbbd", "bb"},
	}

	for k, v := range tc {
		// algo func
		out := longestPalindrome(v.in)

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
func longestPalindrome(s string) string {
	// 20200805
	// 1、动态规划法
	// https://leetcode-cn.com/problems/longest-palindromic-substring/solution/zui-chang-hui-wen-zi-chuan-by-leetcode-solution/
	// Runtime:52 ms, faster than 50.68% of Go online submissions.
	// Memory Usage:6.5 MB, less than 48.25% of Go online submissions.
	n := len(s)
	res := ""
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = true
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = true
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && l+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}

	return res
	// 2、中心扩散法
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
// 示例 1：
//
// 输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//
//
// 示例 2：
//
// 输入: "cbbd"
//输出: "bb"
//
// Related Topics 字符串 动态规划
// 👍 2518 👎 0
