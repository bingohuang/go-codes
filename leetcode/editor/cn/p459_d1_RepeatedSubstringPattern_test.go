// github.com/bingohuang/go-codes
/**
题目: 459. 重复的子字符串
难度: 1
地址: https://leetcode-cn.com/problems/repeated-substring-pattern/
日期：2020-08-24 19:47:56
*/
package test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// input and ouput
type IO459 struct {
	in  string
	out bool
}

func Test459(t *testing.T) {
	// add test cases
	tc := map[string]IO459{
		//"0": {"", false},
		"1": {"abab", true},
		"2": {"aba", false},
	}

	for k, v := range tc {
		// algo func
		out := repeatedSubstringPattern(v.in)

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
func repeatedSubstringPattern(s string) bool {
	// 2020-08-24 19:49 @bingohuang
	// 算法：1、暴力法：两次遍历，一遍子串，一遍全字符串
	// 复杂度：O(n^2)/O(1)
	// 效率：执行耗时:1328 ms,击败了6.17% 的Go用户
	//			内存消耗:8.8 MB,击败了9.09% 的Go用户
	/*n := len(s)
	for i:=1; i <= n/2; i ++ {
		if n % i != 0 {
			continue
		}
		subs := s[0:i]
		ns := ""
		for j := 0; j < n /i ; j ++ {
			ns += subs
		}
		if ns == s {
			return true
		}
	}
	return false*/
	// 2020-08-24 20:06 @bingohuang
	// 算法：字符串匹配法
	// 复杂度：
	// 效率：执行耗时:8 ms,击败了48.15% 的Go用户
	//			内存消耗:6 MB,击败了90.91% 的Go用户
	ss := s + s
	index := strings.Index(ss[1:], s)
	return index != len(s)-1
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
//
// 示例 1:
//
//
//输入: "abab"
//
//输出: True
//
//解释: 可由子字符串 "ab" 重复两次构成。
//
//
// 示例 2:
//
//
//输入: "aba"
//
//输出: False
//
//
// 示例 3:
//
//
//输入: "abcabcabcabc"
//
//输出: True
//
//解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
//
// Related Topics 字符串
// 👍 311 👎 0
