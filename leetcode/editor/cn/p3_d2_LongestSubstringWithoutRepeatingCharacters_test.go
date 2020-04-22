// github.com/bingohuang/go-codes
/**
题目: 3. 无重复字符的最长子串
难度: 2
地址: https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
*/
package test

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// input and ouput
type IO3 struct {
	in  string
	out int
}

func Test3(t *testing.T) {
	// add test cases
	tc := map[string]IO3{
		"1": {"abcabcbb", 3},
		"2": {"bbbbb", 1},
		"3": {"pwwkew", 3},
	}

	for k, v := range tc {
		// algo func
		out := lengthOfLongestSubstring(v.in)

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
func lengthOfLongestSubstring(s string) int {
	// 20200422
	// 解法1：暴力法，找出所有子串，从大到小看看子串是否有重复字符
	// 解法2：线性法
	// 执行耗时:16 ms,击败了28.94% 的Go用户
	// 内存消耗:2.6 MB,击败了80.65% 的Go用户
	// 定义一个哈希集合 set，初始化结果 max 为 0
	set := make(map[byte]struct{})
	max := 0
	// 用快慢指针 i 和 j 扫描一遍字符串，
	// 如果快指针所指向的字符已经出现在哈希集合里，
	// 不断地尝试将慢指针所指向的字符从哈希集合里删除
	for i, j := 0, 0; j < len(s); j++ {
		_, ok := set[s[j]]
		for ok {
			delete(set, s[i])
			i++
			_, ok = set[s[j]]
		}
		// 当快指针的字符加入到哈希集合后，更新一下结果 max
		set[s[j]] = struct{}{}
		max = int(math.Max(float64(max), float64(len(set))))
	}

	return max
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 示例 1:
//
// 输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
// 输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
// 输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
// Related Topics 哈希表 双指针 字符串 Sliding Window
