// github.com/bingohuang/go-codes
/**
题目: 8. 字符串转换整数 (atoi)
难度: 2
地址: https://leetcode-cn.com/problems/string-to-integer-atoi/
*/
package test

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"
)

// input and ouput
type IO8 struct {
	in  string
	out int
}

func Test8(t *testing.T) {
	// add test cases
	tc := map[string]IO8{
		"0": IO8{"", 0},
		"1": IO8{"42", 42},
		"2": IO8{"   -42", -42},
		"3": IO8{"4193 with words", 4193},
		"4": IO8{"-91283472332", -2147483648},
	}

	for k, v := range tc {
		// algo func
		out := myAtoi(v.in)

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
func myAtoi(str string) int {
	// 20200805
	// 1、Go标准库中的strcov.Atoi并不能解这道题
	/*res, _ := strconv.Atoi(str)
	return res*/

	// 2、https://leetcode-cn.com/problems/string-to-integer-atoi/solution/gojie-fa-liang-bu-zou-jian-dan-di-luo-ji-chu-li-zh/
	// Runtime:4 ms, faster than 54.68% of Go online submissions.
	// Memory Usage:2.3 MB, less than 75.58% of Go online submissions.
	// 先处理一些非法数值，留下存数字
	clean := func(s string) (sign int, abs string) {
		s = strings.TrimSpace(s)
		if s == "" {
			return
		}
		switch s[0] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			sign, abs = 1, s
		case '+':
			sign, abs = 1, s[1:]
		case '-':
			sign, abs = -1, s[1:]
		default:
			abs = ""
			return
		}
		for i, b := range abs {
			if b < '0' || '9' < b {
				abs = abs[:i]
				break
			}
		}
		return
	}
	convert := func(sign int, absStr string) int {
		absNum := 0
		for _, b := range absStr {
			absNum = absNum*10 + int(b-'0')
			switch {
			case sign == 1 && absNum > math.MaxInt32:
				return math.MaxInt32
			case sign == -1 && absNum*sign < math.MinInt32:
				return math.MinInt32
			}
		}
		return sign * absNum
	}
	return convert(clean(str))

	// 3、https://leetcode-cn.com/problems/string-to-integer-atoi/solution/zi-fu-chuan-zhuan-huan-zheng-shu-atoi-by-leetcode-/
	// 状态机法
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//请你来实现一个 atoi 函数，使其能将字符串转换成整数。
//
// 首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。接下来的转化规则如下：
//
//
// 如果第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字字符组合起来，形成一个有符号整数。
// 假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成一个整数。
// 该字符串在有效的整数部分之后也可能会存在多余的字符，那么这些字符可以被忽略，它们对函数不应该造成影响。
//
//
// 注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换，即无法进行有效转换。
//
// 在任何情况下，若函数不能进行有效的转换时，请返回 0 。
//
// 提示：
//
//
// 本题中的空白字符只包括空格字符 ' ' 。
// 假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231, 231 − 1]。如果数值超过这个范围，请返回 INT_MAX (231
// − 1) 或 INT_MIN (−231) 。
//
//
//
//
// 示例 1:
//
// 输入: "42"
//输出: 42
//
//
// 示例 2:
//
// 输入: "   -42"
//输出: -42
//解释: 第一个非空白字符为 '-', 它是一个负号。
//     我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。
//
//
// 示例 3:
//
// 输入: "4193 with words"
//输出: 4193
//解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字。
//
//
// 示例 4:
//
// 输入: "words and 987"
//输出: 0
//解释: 第一个非空字符是 'w', 但它不是数字或正、负号。
//     因此无法执行有效的转换。
//
// 示例 5:
//
// 输入: "-91283472332"
//输出: -2147483648
//解释: 数字 "-91283472332" 超过 32 位有符号整数范围。
//     因此返回 INT_MIN (−231) 。
//
// Related Topics 数学 字符串
// 👍 781 👎 0
