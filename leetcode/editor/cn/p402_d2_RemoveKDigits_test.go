// github.com/bingohuang/go-codes
/**
题目: 402. 移掉K位数字
难度: 2
地址: https://leetcode-cn.com/problems/remove-k-digits/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO402 struct {
	in1 string
	in2 int
	out string
}

func Test402(t *testing.T) {
	// add test cases
	tc := map[string]IO402{
		"1": IO402{"1432219", 3, "1219"},
		"2": IO402{"10200", 1, "200"},
		"3": IO402{"10", 2, "0"},
	}

	for k, v := range tc {
		// algo func
		out := removeKdigits(v.in1, v.in2)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput: %v|%v\n", v.in1, v.in2)
		fmt.Printf("\toutput: %v\n", out)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func removeKdigits(num string, k int) string {
	var stack []uint8               // 使用 slice 作为栈
	var result string               // 存储最终结果字符串
	for i := 0; i < len(num); i++ { // 从最高位扫描数字字符串num
		number := num[i] - '0' // 将字符串转化为整数
		// 条件：当栈不为空，且栈顶元素大于下一个数字，且 k 还大于0
		for len(stack) != 0 && stack[len(stack)-1] > number && k > 0 {
			stack = stack[:len(stack)-1] // 弹出栈顶元素
			k--                          // 剔除一个数字，k 减一
		}
		// 条件：如果数字不为空，或者数字为空但是栈不为空
		if number != 0 || len(stack) != 0 {
			stack = append(stack, number) // 入栈
		}
	}
	// 条件：如果栈不空且还能删除数字
	for len(stack) != 0 && k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	// 将栈中元素转化为字符串
	for _, v := range stack {
		result += string('0' + v)
	}
	// 注意为空情况
	if result == "" {
		return "0"
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
//
// 注意:
//
//
// num 的长度小于 10002 且 ≥ k。
// num 不会包含任何前导零。
//
//
// 示例 1 :
//
//
//输入: num = "1432219", k = 3
//输出: "1219"
//解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
//
//
// 示例 2 :
//
//
//输入: num = "10200", k = 1
//输出: "200"
//解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
//
//
// 示例 3 :
//
//
//输入: num = "10", k = 2
//输出: "0"
//解释: 从原数字移除所有的数字，剩余为空就是0。
//
// Related Topics 栈 贪心算法
