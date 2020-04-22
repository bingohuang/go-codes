// github.com/bingohuang/go-codes
/**
题目: 516. 最长回文子序列
难度: 2
地址: https://leetcode-cn.com/problems/longest-palindromic-subsequence/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO516 struct {
	in  string
	out int
}

func Test516(t *testing.T) {
	// add test cases
	tc := map[string]IO516{
		"0": {"", 0},
		"1": {"a", 1},
		"2": {"aab", 2},
		"3": {"cbbd", 2},
		"4": {"bbbab", 4},
	}

	for k, v := range tc {
		// algo func
		out := longestPalindromeSubseq(v.in)

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
func longestPalindromeSubseq(s string) int {

	// 20200422:
	// 执行耗时:44 ms,击败了31.39% 的Go用户
	// 内存消耗:17 MB,击败了100.00% 的Go用户
	// 思路：
	// 对于回文来说，必须保证两头的字符都相同
	// 用 d[i][j]表示从字符串第i个字符到第j个字符之间的最长回文
	// 比较这段区间外的两个字符，如果它们相等，肯定能构成新的最长回文
	// 最终的最长回文会保存在dp[0][n-1]
	// 递推公式如下：
	// 当首位两个元素相等时：dp[0][n-1]=dp[1][n-2]+2
	// 否则：dp[0][n-1]=max(dp[1][n-1], dp[0][n-2])

	n := len(s)
	if n <= 1 {
		return n
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for l := 2; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				dp[i][j] = 2
				if l != 2 {
					dp[i][j] += dp[i+1][j-1]
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}

//func updateDP(dp [][]int, s string, i, j int)  {
//
//}

//func max(x, y int) int {
//	if x < y {
//		return y
//	}
//	return x
//}
//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个字符串s，找到其中最长的回文子序列。可以假设s的最大长度为1000。
//
// 示例 1:
//输入:
//
//
//"bbbab"
//
//
// 输出:
//
//
//4
//
//
// 一个可能的最长回文子序列为 "bbbb"。
//
// 示例 2:
//输入:
//
//
//"cbbd"
//
//
// 输出:
//
//
//2
//
//
// 一个可能的最长回文子序列为 "bb"。
// Related Topics 动态规划
