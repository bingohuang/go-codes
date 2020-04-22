// github.com/bingohuang/go-codes
/**
题目: 70. 爬楼梯
难度: 1
地址: https://leetcode-cn.com/problems/climbing-stairs/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO70 struct {
	in  int
	out int
}

func Test70(t *testing.T) {
	// add test cases
	tc := map[string]IO70{
		//"1": {1, 1},
		"2": {2, 2},
		//"3": {3, 3},
		//"4": {4, 5},
	}

	for k, v := range tc {
		// algo func
		out := climbStairs(v.in)

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
func climbStairs(n int) int {
	// 20200422
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2 MB,击败了63.64% 的Go用户
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]

	// 20180630
	// runtime:0 ms
	// memory:N/A
	//if in <= 0 {
	//	return 0
	//}
	//
	//if in == 1 || in == 2 {
	//	return in
	//}
	//
	//s1 := 1
	//s2 := 2
	//var s3 int
	//
	//for s := 3; s <= in; s++ {
	//	s3 = s1 + s2
	//	s1 = s2
	//	s2 = s3
	//}
	//
	//return s3
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 注意：给定 n 是一个正整数。
//
// 示例 1：
//
// 输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//
// 示例 2：
//
// 输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
//
// Related Topics 动态规划
