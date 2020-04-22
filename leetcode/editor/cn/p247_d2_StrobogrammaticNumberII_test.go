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
type IO247 struct {
	in  int
	out []string
}

func Test247(t *testing.T) {
	// add test cases
	tc := map[string]IO247{
		//"0": {0, []string{""}},
		//"1": {1, []string{"0", "1", "8"}},
		//"2": {2, []string{"11", "69", "88", "96"}},
		//"3": {3, []string{"101", "609", "808", "906", "111", "619", "818", "916", "181", "689", "888", "986"}},
		//"4": {4, []string{}},
		"5": {5, []string{}},
	}

	for k, v := range tc {
		// algo func
		out := findStrobogrammatic(v.in)

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
func findStrobogrammatic(n int) []string {
	if n < 0 {
		return nil
	}

	return helper(n, n)
}

func helper(n, m int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"0", "1", "8"}
	}
	list := helper(n-2, m)
	var res []string
	for i := 0; i < len(list); i++ {
		if n != m {
			res = append(res, "0"+list[i]+"0")
		}
		res = append(res, "1"+list[i]+"1")
		res = append(res, "6"+list[i]+"9")
		res = append(res, "8"+list[i]+"8")
		res = append(res, "9"+list[i]+"6")
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
// 中心对称数是指一个数字在旋转了 180 度之后看起来依旧相同的数字（或者上下颠倒地看）。
//
// 找到所有长度为 n 的中心对称数。
//
// 示例 :
//
// 输入:  n = 2
// 输出: ["11","69","88","96"]
