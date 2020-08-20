// github.com/bingohuang/go-codes
/**
题目: 125. 验证回文串
难度: 1
地址: https://leetcode-cn.com/problems/valid-palindrome/
日期：2020-08-20 11:07:48
*/
package test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// input and ouput
type IO125 struct {
	in  string
	out bool
}

func Test125(t *testing.T) {
	// add test cases
	tc := map[string]IO125{
		//"0": {"", true},
		//"1": {"A man, a plan, a canal: Panama", true},
		//"2": {"race a car", false},
		"3": {"ab2a", false},
	}

	for k, v := range tc {
		// algo func
		out := isPalindrome(v.in)

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
func isPalindrome(s string) bool {
	// 2020-08-20 11:27 @bingohuang
	// 算法：1、双指针+标准库
	// 复杂度：O(N)/O(1)
	// 效率：执行耗时:0 ms,击败了100.00% 的Go用户
	//			内存消耗:2.9 MB,击败了78.33% 的Go用户
	l := 0
	r := len(s) - 1
	s = strings.ToLower(s)
	for l < r {
		for l < r {
			if ('0' <= s[l] && s[l] <= '9') || ('a' <= s[l] && s[l] <= 'z') {
				break
			}
			l++
		}
		for l < r {
			if ('0' <= s[r] && s[r] <= '9') || ('a' <= s[r] && s[r] <= 'z') {
				break
			}
			r--
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
	// 2020-08-20 11:32 @bingohuang
	// 算法：2、数组复制，再通过普通回文串来判断 s[i] == s[n-i-1]
	// 复杂度：
	// 效率：

}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
// 说明：本题中，我们将空字符串定义为有效的回文串。
//
// 示例 1:
//
// 输入: "A man, a plan, a canal: Panama"
//输出: true
//
//
// 示例 2:
//
// 输入: "race a car"
//输出: false
//
// Related Topics 双指针 字符串
// 👍 264 👎 0
