// github.com/bingohuang/go-codes
/**
题目: 131. 分割回文串
难度: 2
地址: https://leetcode-cn.com/problems/palindrome-partitioning/
日期：2020-08-20 15:34:50
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO131 struct {
	in  string
	out [][]string
}

func Test131(t *testing.T) {
	// add test cases
	// 测试用例:"cbbbcc"
	// 测试结果:[["c","b","b","b","cc","c"],["c","b","b","b","cc"],["c","b","bb","c","c"],["c","b","bb","cc"],["c","bb","b","c","c"],["c","bb","b","cc"],["c","bbb","cc","c"],["c","bbb","cc"],["cbbbc","c"]]
	// 期望结果:[["c","b","b","b","c","c"],["c","b","b","b","cc"],["c","b","bb","c","c"],["c","b","bb","cc"],["c","bb","b","c","c"],["c","bb","b","cc"],["c","bbb","c","c"],["c","bbb","cc"],["cbbbc","c"]]
	tc := map[string]IO131{
		//"0": {"aab", [][]string{
		//	 {"aa","b"},
		//	 {"a","a","b"},
		//}},
		"1": {"cbbbcc", [][]string{
			{"c", "b", "b", "b", "c", "c"}, {"c", "b", "b", "b", "cc"}, {"c", "b", "bb", "c", "c"}, {"c", "b", "bb", "cc"}, {"c", "bb", "b", "c", "c"}, {"c", "bb", "b", "cc"}, {"c", "bbb", "c", "c"}, {"c", "bbb", "cc"}, {"cbbbc", "c"},
		}},
	}

	for k, v := range tc {
		// algo func
		out := partition(v.in)

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
func partition(s string) [][]string {
	// 2020-08-20 15:43 @bingohuang
	// 算法：1、回溯法：https://leetcode-cn.com/problems/palindrome-partitioning/solution/hui-su-you-hua-jia-liao-dong-tai-gui-hua-by-liweiw/
	// 复杂度：
	// 效率：执行耗时:4 ms,击败了100.00% 的Go用户
	//			内存消耗:5.3 MB,击败了85.25% 的Go用户
	n := len(s)
	if n == 0 {
		return nil
	}
	var res [][]string
	var stack []string

	// 检查s[left:right+1]是否为回文串
	check := func(s string, left, right int) bool {
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}
	// 回溯算法
	var backtrace func(string, int, int, []string)
	backtrace = func(s string, start int, length int, path []string) {
		if start == length {
			// 此处一定要用tmp+copy的方式
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := start; i < length; i++ {
			if !check(s, start, i) {
				continue
			}
			path = append(path, s[start:i+1])
			backtrace(s, i+1, length, path)
			path = path[:len(path)-1]
		}
	}
	backtrace(s, 0, n, stack)

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
//
// 返回 s 所有可能的分割方案。
//
// 示例:
//
// 输入: "aab"
//输出:
//[
//  ["aa","b"],
//  ["a","a","b"]
//]
// Related Topics 回溯算法
// 👍 345 👎 0
