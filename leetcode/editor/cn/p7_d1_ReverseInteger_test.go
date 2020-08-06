// github.com/bingohuang/go-codes
/**
题目: 7. 整数反转
难度: 1
地址: https://leetcode-cn.com/problems/reverse-integer/
*/
package test

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// input and ouput
type IO7 struct {
	in  int
	out int
}

func Test7(t *testing.T) {
	// add test cases
	// 示例 1:
	//
	// 输入: 123
	//输出: 321
	//
	//
	// 示例 2:
	//
	// 输入: -123
	//输出: -321
	//
	//
	// 示例 3:
	//
	// 输入: 120
	//输出: 21
	tc := map[string]IO7{
		"0":  {0, 0},
		"1":  {123, 321},
		"2":  {-123, -321},
		"3":  {120, 21},
		"4*": {1534236469, 0}, // not 9646324351
	}

	for k, v := range tc {
		// algo func
		out := reverse(v.in)

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
func reverse(x int) int {
	// 20200806
	// 1. 标准库法：注意超过int32的用例
	// 执行耗时:4 ms,击败了48.10% 的Go用户
	// 内存消耗:2.2 MB,击败了5.37% 的Go用户
	/*sign := 1
	if x < 0 {
		sign = -1
		x = x * sign
	}
	s := strconv.Itoa(x)
	s = strings.Trim(s, "0")

	ns := ""
	for _, v := range s {
		ns = string(v) + ns
	}
	res, _ := strconv.Atoi(ns)
	if res > math.MaxInt32 {
		return 0
	}
	return res * sign*/

	// 20200806
	// 2. 数学法
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.2 MB,击败了63.90% 的Go用户
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > math.MaxInt32%10) {
			rev = 0
			return rev
		} else if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < math.MinInt32%10) {
			rev = 0
			return rev
		}
		rev = rev*10 + pop
	}
	return rev
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
// 示例 1:
//
// 输入: 123
//输出: 321
//
//
// 示例 2:
//
// 输入: -123
//输出: -321
//
//
// 示例 3:
//
// 输入: 120
//输出: 21
//
//
// 注意:
//
// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231, 231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
// Related Topics 数学
// 👍 2088 👎 0
