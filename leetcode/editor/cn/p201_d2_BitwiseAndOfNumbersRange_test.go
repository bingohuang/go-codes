// github.com/bingohuang/go-codes
/**
题目: 201. 数字范围按位与
难度: 2
地址: https://leetcode-cn.com/problems/bitwise-and-of-numbers-range/
日期：2020-08-24 20:22:27
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO201 struct {
	in1 int
	in2 int
	out int
}

func Test201(t *testing.T) {
	// add test cases
	tc := map[string]IO201{
		"0": {0, 0, 0},
		"1": {5, 7, 4},
		"2": {0, 1, 0},
	}

	for k, v := range tc {
		// algo func
		out := rangeBitwiseAnd(v.in1, v.in2)

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
func rangeBitwiseAnd(m int, n int) int {
	// 2020-08-24 20:25 @bingohuang
	// 算法：位移法
	// 复杂度：O(logn)/O(1)
	// 效率：执行耗时:16 ms,击败了69.79% 的Go用户
	//			内存消耗:6.1 MB,击败了46.15% 的Go用户
	shift := 0
	for m < n {
		m, n = m>>1, n>>1
		shift++
	}
	return m << shift
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。
//
// 示例 1:
//
// 输入: [5,7]
//输出: 4
//
// 示例 2:
//
// 输入: [0,1]
//输出: 0
// Related Topics 位运算
// 👍 212 👎 0
