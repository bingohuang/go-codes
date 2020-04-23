// github.com/bingohuang/go-codes
/**
题目: 316. 去除重复字母
难度: 3
地址: https://leetcode-cn.com/problems/remove-duplicate-letters/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO316 struct {
	in  string
	out string
}

func Test316(t *testing.T) {
	// add test cases
	tc := map[string]IO316{
		//"1": {"bcabc", "abc"},
		//"2": {"cbacdcbc", "acdb"},
		"3": {"bbcaac", "bac"},
	}

	for k, v := range tc {
		// algo func
		out := removeDuplicateLetters(v.in)

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
func removeDuplicateLetters(s string) string {
	// 20200423
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.2 MB,击败了100.00% 的Go用户
	// 执行用时 : 4 ms
	//内存消耗 : 2.2 MB
	stack := make([]rune, 0)
	seen := make(map[rune]bool)
	occur := make(map[rune]int)
	for _, v := range s {
		if _, ok := occur[v]; !ok {
			occur[v] = 1
		} else {
			occur[v]++
		}
	}
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			for len(stack) > 0 && v < stack[len(stack)-1] {
				e := stack[len(stack)-1]
				if o, ok := occur[e]; ok && o > 1 {
					stack = stack[:len(stack)-1]
					occur[e]--
					delete(seen, e)
				} else {
					break
				}
			}
			stack = append(stack, v)
			seen[v] = true
		} else {
			occur[v]--
		}
	}
	res := ""
	for _, v := range stack {
		res += string(v)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给你一个仅包含小写字母的字符串，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
//
//
//
// 示例 1:
//
// 输入: "bcabc"
//输出: "abc"
//
//
// 示例 2:
//
// 输入: "cbacdcbc"
//输出: "acdb"
//
//
//
// 注意：该题与 1081 https://leetcode-cn.com/problems/smallest-subsequence-of-distinct
//-characters 相同
// Related Topics 栈 贪心算法
