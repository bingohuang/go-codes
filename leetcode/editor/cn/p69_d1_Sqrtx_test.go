// github.com/bingohuang/go-codes
/**
题目: 69. x 的平方根
难度: 1
地址: https://leetcode-cn.com/problems/sqrtx/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO69 struct {
	in  int
	out int
}

func Test69(t *testing.T) {
	// add test cases
	tc := map[string]IO69{
		"0": {1, 1},
		"1": {2, 1},
		"2": {4, 2},
		"3": {8, 2},
	}

	for k, v := range tc {
		// algo func
		out := mySqrt(v.in)

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
func mySqrt(x int) int {
	// 20200818
	// 1、暴力法
	// 执行耗时:20 ms,击败了6.35% 的Go用户
	// 内存消耗:2.2 MB,击败了100.00% 的Go用户
	/*for i := 1; i <= x; i ++ {
		if i * i == x {
			return i
		} else if i * i > x {
			return i - 1
		}
	}
	return 0*/
	// 2、二分法
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.2 MB,击败了58.22% 的Go用户
	l, r := 0, x
	res := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//实现 int sqrt(int x) 函数。
//
// 计算并返回 x 的平方根，其中 x 是非负整数。
//
// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
//
// 示例 1:
//
// 输入: 4
//输出: 2
//
//
// 示例 2:
//
// 输入: 8
//输出: 2
//说明: 8 的平方根是 2.82842...,
//     由于返回类型是整数，小数部分将被舍去。
//
// Related Topics 数学 二分查找
// 👍 480 👎 0
