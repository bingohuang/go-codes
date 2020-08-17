// github.com/bingohuang/go-codes
/**
题目: 28. 实现 strStr()
难度: 1
地址: https://leetcode-cn.com/problems/implement-strstr/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO28 struct {
	in1 string
	in2 string
	out int
}

func Test28(t *testing.T) {
	// add test cases
	tc := map[string]IO28{
		"0": {"", "", 0},
		"1": {"hello", "ll", 2},
		"2": {"aaaaa", "bba", -1},
	}

	for k, v := range tc {
		// algo func
		out := strStr(v.in1, v.in2)

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
func strStr(haystack string, needle string) int {
	// 20200817
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}
	// 1、标准库法
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.3 MB,击败了64.65% 的Go用户
	/*
		return strings.Index(haystack, needle)*/
	// 2、滑动窗口法
	// O((N-L)*L)/O(1)
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.3 MB,击败了100.00% 的Go用户
	n, l := len(haystack), len(needle)
	for start := 0; start < n-l+1; start++ {
		if haystack[start:start+l] == needle {
			return start
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//实现 strStr() 函数。
//
// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如
//果不存在，则返回 -1。
//
// 示例 1:
//
// 输入: haystack = "hello", needle = "ll"
//输出: 2
//
//
// 示例 2:
//
// 输入: haystack = "aaaaa", needle = "bba"
//输出: -1
//
//
// 说明:
//
// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//
// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
// Related Topics 双指针 字符串
// 👍 533 👎 0
