// github.com/bingohuang/go-codes
/**
题目: 60. 第k个排列
难度: 2
地址: https://leetcode-cn.com/problems/permutation-sequence/
日期：2020-09-05 11:25:55
*/
package test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

// input and ouput
type IO60 struct {
	in1 int
	in2 int
	out string
}

func Test60(t *testing.T) {
	// add test cases
	tc := map[string]IO60{
		"0": {1, 1, "1"},
	}

	for k, v := range tc {
		// algo func
		out := getPermutation(v.in1, v.in2)

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
func getPermutation(n int, k int) string {
	// 2020-09-05 11:30 @bingohuang
	// 算法：1、数学法：要理解背后的数学公式和思路
	// 复杂度：O(n^2)/O(n）
	// 效率：执行耗时:0 ms,击败了100.00% 的Go用户
	//			内存消耗:2.1 MB,击败了26.88% 的Go用户
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}
	k--

	ans := ""
	valid := make([]int, n+1)
	for i := 0; i < len(valid); i++ {
		valid[i] = 1
	}
	for i := 1; i <= n; i++ {
		order := k/factorial[n-i] + 1
		for j := 1; j <= n; j++ {
			order -= valid[j]
			if order == 0 {
				ans += strconv.Itoa(j)
				valid[j] = 0
				break
			}
		}
		k %= factorial[n-i]
	}
	return ans

	// 2020-09-05 11:31 @bingohuang
	// 算法：2、回溯法：TODO
	// 复杂度：
	// 效率：

}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
//
// 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
//
//
// "123"
// "132"
// "213"
// "231"
// "312"
// "321"
//
//
// 给定 n 和 k，返回第 k 个排列。
//
// 说明：
//
//
// 给定 n 的范围是 [1, 9]。
// 给定 k 的范围是[1, n!]。
//
//
// 示例 1:
//
// 输入: n = 3, k = 3
//输出: "213"
//
//
// 示例 2:
//
// 输入: n = 4, k = 9
//输出: "2314"
//
// Related Topics 数学 回溯算法
// 👍 338 👎 0
