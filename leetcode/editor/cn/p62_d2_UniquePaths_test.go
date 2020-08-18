// github.com/bingohuang/go-codes
/**
题目: 62. 不同路径
难度: 2
地址: https://leetcode-cn.com/problems/unique-paths/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO62 struct {
	in1 int
	in2 int
	out int
}

func Test62(t *testing.T) {
	// add test cases
	tc := map[string]IO62{
		//"0": {0, 0,0},
		//"1": {1, 1,1},
		"2": {2, 3, 3},
		"3": {7, 3, 28},
	}

	for k, v := range tc {
		// algo func
		out := uniquePaths(v.in1, v.in2)

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
func uniquePaths(m int, n int) int {
	// 20200818
	// 1、数学法
	// 一共m行n列，其中需要向下走m-1步，向右走n-1步，一共走m+n-2步，所以就是在m+n-2步中选出哪m-1步是向下走的，其余自动向右走
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:1.9 MB,击败了100.00% 的Go用户
	/*if m == 1 || n == 1 {
		return 1
	}
	if m > n {
		m, n = n, m
	}
	tmp, res := 1, 1
	for i := 1; i <= m-1; i++ {
		tmp *= i
	}
	for i := n; i <= m+n-2; i++ {
		res *= i
	}
	return res / tmp*/
	// 2、动态规划法
	// dp[i][j]=1 i=0 or j=0
	// dp[i][j]=dp[i-1][j]+dp[i][j-1]
	// O(m*n)/O(m*n)
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2 MB,击败了50.54% 的Go用户
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
// 问总共有多少条不同的路径？
//
//
//
// 例如，上图是一个7 x 3 的网格。有多少可能的路径？
//
//
//
// 示例 1:
//
// 输入: m = 3, n = 2
//输出: 3
//解释:
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向右 -> 向下
//2. 向右 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向右
//
//
// 示例 2:
//
// 输入: m = 7, n = 3
//输出: 28
//
//
//
// 提示：
//
//
// 1 <= m, n <= 100
// 题目数据保证答案小于等于 2 * 10 ^ 9
//
// Related Topics 数组 动态规划
// 👍 646 👎 0
