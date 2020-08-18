// github.com/bingohuang/go-codes
/**
题目: 66. 加一
难度: 1
地址: https://leetcode-cn.com/problems/plus-one/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO66 struct {
	in  []int
	out []int
}

func Test66(t *testing.T) {
	// add test cases
	tc := map[string]IO66{
		"0": {[]int{0}, []int{1}},
		"1": {[]int{1, 2, 3}, []int{1, 2, 4}},
		"2": {[]int{9, 9}, []int{1, 0, 0}},
		"3": {[]int{9, 1}, []int{9, 2}},
	}

	for k, v := range tc {
		// algo func
		out := plusOne(v.in)

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
func plusOne(digits []int) []int {
	// 20200818
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.3 MB,击败了6.98% 的Go用户
	/*var res []int
	n := len(digits)
	// plus one
	plus := digits[n-1] + 1
	carry := plus/10
	if carry == 1 {
		res = append(res, 0)
	} else {
		res = append(res, plus)
	}

	for i:=n-2;i>=0;i--{
		plus = digits[i] + carry
		carry = plus/10
		if carry == 1 {
			res = append([]int{0}, res...)
		} else {
			res = append([]int{plus}, res...)
		}

	}
	if carry == 1 {
		res = append([]int{carry}, res...)
	}
	return res*/

	// https://leetcode-cn.com/problems/plus-one/solution/gojie-fa-xiang-xi-zhu-shi-kao-lu-jin-wei-yu-shu-zu/
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2 MB,击败了64.44% 的Go用户
	/*for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
			continue
		} else {
			digits[i]++
			break
		}
	}
	return digits*/

	// https://leetcode-cn.com/problems/plus-one/solution/java-shu-xue-jie-ti-by-yhhzw/
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.1 MB,击败了6.98% 的Go用户
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//
// 你可以假设除了整数 0 之外，这个整数不会以零开头。
//
// 示例 1:
//
// 输入: [1,2,3]
//输出: [1,2,4]
//解释: 输入数组表示数字 123。
//
//
// 示例 2:
//
// 输入: [4,3,2,1]
//输出: [4,3,2,2]
//解释: 输入数组表示数字 4321。
//
// Related Topics 数组
// 👍 526 👎 0
