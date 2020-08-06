// github.com/bingohuang/go-codes
/**
题目: 13. 罗马数字转整数
难度: 1
地址: https://leetcode-cn.com/problems/roman-to-integer/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO13 struct {
	in  string
	out int
}

func Test13(t *testing.T) {
	// add test cases
	// 示例 1:
	//
	// 输入: "III"
	//输出: 3
	//
	// 示例 2:
	//
	// 输入: "IV"
	//输出: 4
	//
	// 示例 3:
	//
	// 输入: "IX"
	//输出: 9
	//
	// 示例 4:
	//
	// 输入: "LVIII"
	//输出: 58
	//解释: L = 50, V= 5, III = 3.
	//
	//
	// 示例 5:
	//
	// 输入: "MCMXCIV"
	//输出: 1994
	//解释: M = 1000, CM = 900, XC = 90, IV = 4.
	tc := map[string]IO13{
		"0":  {"", 0},
		"1":  {"III", 3},
		"2":  {"IV", 4},
		"3":  {"IX", 9},
		"4":  {"LVIII", 58},
		"5":  {"MCMXCIV", 1994},
		"6*": {"MMMXLV", 3045},
	}

	for k, v := range tc {
		// algo func
		out := romanToInt(v.in)

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
func romanToInt(s string) int {
	// 20200806
	// 1、从左往右遍历
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:3.1 MB,击败了100.00% 的Go用户
	/*ans := 0
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] == 'I' && i+1 < l && s[i+1] == 'V' {
			ans += 4
			i ++
		} else if s[i] == 'I' && i+1 < l && s[i+1] == 'X' {
			ans += 9
			i ++
		} else if s[i] == 'X' && i+1 < l && s[i+1] == 'L' {
			ans += 40
			i ++
		} else if s[i] == 'X' && i+1 < l && s[i+1] == 'C' {
			ans += 90
			i ++
		} else if s[i] == 'C' && i+1 < l && s[i+1] == 'D' {
			ans += 400
			i ++
		} else if s[i] == 'C' && i+1 < l && s[i+1] == 'M' {
			ans += 900
			i ++
		} else {
			switch s[i] {
			case 'I':
				ans += 1
			case 'V':
				ans += 5
			case 'X':
				ans += 10
			case 'L':
				ans += 50
			case 'C':
				ans += 100
			case 'D':
				ans += 500
			case 'M':
				ans += 1000
			}
		}
	}
	return ans*/

	// 20200806
	// 2、从右往左遍历，可以使代码更简洁 - 结合map
	// 执行耗时:4 ms,击败了94.70% 的Go用户
	// 内存消耗:3.1 MB,击败了60.00% 的Go用户
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ans := 0
	pre := 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := m[s[i]]
		if cur >= pre {
			ans += cur
		} else {
			ans -= cur
		}
		pre = cur
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
//
// 字符          数值
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//
// 例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做 XXVII, 即为 XX + V + I
//I 。
//
// 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5
// 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
//
//
// I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
// X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
// C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
//
//
// 给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。
//
// 示例 1:
//
// 输入: "III"
//输出: 3
//
// 示例 2:
//
// 输入: "IV"
//输出: 4
//
// 示例 3:
//
// 输入: "IX"
//输出: 9
//
// 示例 4:
//
// 输入: "LVIII"
//输出: 58
//解释: L = 50, V= 5, III = 3.
//
//
// 示例 5:
//
// 输入: "MCMXCIV"
//输出: 1994
//解释: M = 1000, CM = 900, XC = 90, IV = 4.
// Related Topics 数学 字符串
// 👍 981 👎 0
