// github.com/bingohuang/go-codes
/**
题目: 91. 解码方法
难度: 2
地址: https://leetcode-cn.com/problems/decode-ways/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO91 struct {
	in  string
	out int
}

func Test91(t *testing.T) {
	// add test cases
	tc := map[string]IO91{
		//"1": {"1", 1},
		//"2": {"12", 2},
		//"3": {"226", 3},
		//"4": {"20", 1},
		//"5": {"021", 0},
		"6": {"220", 1},
		"7": {"230", 0},
	}

	for k, v := range tc {
		// algo func
		out := numDecodings(v.in)

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
func numDecodings(s string) int {
	// 20200422:
	// 执行耗时:972 ms,击败了5.22% 的Go用户
	// 内存消耗:2.2 MB,击败了50.00% 的Go用户
	if s[0] == '0' { // 此处必须
		return 0
	}
	return decode(s, len(s)-1)
}

func decode(s string, index int) int {
	if index <= 0 {
		return 1
	}
	count := 0
	curr := s[index]
	prev := s[index-1]

	if curr > '0' { // 如果等于0，则不能进一步
		count = decode(s, index-1)
	}
	if prev == '1' || (prev == '2' && curr <= '6') {
		count += decode(s, index-2)
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//一条包含字母 A-Z 的消息通过以下方式进行了编码：
//
// 'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//
//
// 给定一个只包含数字的非空字符串，请计算解码方法的总数。
//
// 示例 1:
//
// 输入: "12"
//输出: 2
//解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
//
//
// 示例 2:
//
// 输入: "226"
//输出: 3
//解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
//
// Related Topics 字符串 动态规划
