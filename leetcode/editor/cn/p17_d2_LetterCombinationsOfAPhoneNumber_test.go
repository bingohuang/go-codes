// github.com/bingohuang/go-codes
/**
题目: 17. 电话号码的字母组合
难度: 2
地址: https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO17 struct {
	in  string
	out []string
}

func Test17(t *testing.T) {
	// add test cases
	tc := map[string]IO17{
		"0": {"", nil},
		"1": {"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}

	for k, v := range tc {
		// algo func
		out := letterCombinations(v.in)

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
func letterCombinations(digits string) []string {
	// 20200813
	//  O(3^N * 4^M)/O(3^N * 4^M)
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2 MB,击败了90.70% 的Go用户
	var res []string
	m := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var backtrack func(combination string, nextDigits string)
	backtrack = func(combination string, nextDigits string) {
		if nextDigits == "" {
			res = append(res, combination)
		} else {
			digit := nextDigits[:1]
			letter := m[digit]
			for _, l := range letter {
				backtrack(combination+string(l), nextDigits[1:])
			}
		}

	}
	if digits != "" {
		backtrack("", digits)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
// 示例:
//
// 输入："23"
//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//
//
// 说明:
//尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
// Related Topics 字符串 回溯算法
// 👍 825 👎 0
