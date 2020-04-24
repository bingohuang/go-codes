// github.com/bingohuang/go-codes
/**
题目: 179. 最大数
难度: 2
地址: https://leetcode-cn.com/problems/largest-number/
*/
package test

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// input and ouput
type IO179 struct {
	in  []int
	out string
}

func Test179(t *testing.T) {
	// add test cases
	tc := map[string]IO179{
		"1": {[]int{10, 2}, "210"},
		"2": {[]int{3, 30, 34, 5, 9}, "9534330"},
		"3": {[]int{0, 0, 0}, "0"},
	}

	for k, v := range tc {
		// algo func
		out := largestNumber(v.in)

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
func largestNumber(nums []int) string {
	// 20200424
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.3 MB,击败了100.00% 的Go用户
	strs := make([]string, len(nums))
	for i, v := range nums {
		strs[i] = strconv.Itoa(v)
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})
	res := strings.Join(strs, "")
	if res[0] == '0' {
		return "0"
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。
//
// 示例 1:
//
// 输入: [10,2]
//输出: 210
//
// 示例 2:
//
// 输入: [3,30,34,5,9]
//输出: 9534330
//
// 说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
// Related Topics 排序
