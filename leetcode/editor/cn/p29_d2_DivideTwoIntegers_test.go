// github.com/bingohuang/go-codes
/**
题目: 29. 两数相除
难度: 2
地址: https://leetcode-cn.com/problems/divide-two-integers/
*/
package test

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// input and ouput
type IO29 struct {
	in1 int
	in2 int
	out int
}

func Test29(t *testing.T) {
	// add test cases
	tc := map[string]IO29{
		//"0": {0, 1, 0},
		////"1": {1, 0, math.MinInt32},
		//"2": {10, 3, 3},
		//"3": {7, -3, -2},
		//"4": {-7, 3, -2},
		"5": {10000, 3, 3333},
	}

	for k, v := range tc {
		// algo func
		out := divide(v.in1, v.in2)

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
func divide(dividend int, divisor int) int {
	// 20200817
	// 除数不为 0。
	/*if divisor == 0 {
		return math.MaxInt32
	}*/
	// 1、标准除法
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.5 MB,击败了100.00% 的Go用户
	/*res :=  dividend/divisor
	// 如果除法结果溢出，则返回 231 − 1
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res*/
	// 2、采用减法，注意正负数
	// Time Limit Exceeded
	/*positive := 1
	if dividend < 0 {
		dividend = -dividend
		positive = -1 * positive
	}
	if divisor < 0 {
		divisor = - divisor
		positive = -1 * positive
	}
	quotient := 0
	for dividend >= divisor {
		dividend -= divisor
		quotient ++
		if quotient > math.MaxInt32 {
			quotient = math.MaxInt32
			break
		}
	}
	return positive * quotient*/

	// 执行耗时:4 ms,击败了61.45% 的Go用户
	// 内存消耗:2.5 MB,击败了61.67% 的Go用户
	// 采用二分法的思想，dividend每次减去2^n个divisor（尽可能多），同时reslut每次加2^n
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	k := (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0)
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	var ans int
	for dividend >= divisor {
		// 按照两倍的速度增长来除
		var tmp int
		for dividend >= divisor<<tmp {
			tmp++
		}
		dividend -= divisor << (tmp - 1)
		ans += 1 << (tmp - 1)
	}
	if k {
		return ans
	}
	return -ans
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
//
// 返回被除数 dividend 除以除数 divisor 得到的商。
//
// 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
//
//
//
// 示例 1:
//
// 输入: dividend = 10, divisor = 3
//输出: 3
//解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
//
// 示例 2:
//
// 输入: dividend = 7, divisor = -3
//输出: -2
//解释: 7/-3 = truncate(-2.33333..) = -2
//
//
//
// 提示：
//
//
// 被除数和除数均为 32 位有符号整数。
// 除数不为 0。
// 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231, 231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
//
// Related Topics 数学 二分查找
// 👍 399 👎 0
