// github.com/bingohuang/go-codes
/**
题目: 面试题60. n个骰子的点数
难度: 1
地址: https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof/
*/
package test

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// input and ouput
type IO面试题60 struct {
	in  int
	out []float64
}

func Test面试题60(t *testing.T) {
	// add test cases
	tc := map[string]IO面试题60{
		"1": {1, []float64{0.16667, 0.16667, 0.16667, 0.16667, 0.16667, 0.16667}},
		"2": {2, []float64{0.02778, 0.05556, 0.08333, 0.11111, 0.13889, 0.16667, 0.13889, 0.11111, 0.08333, 0.05556, 0.02778}},
	}

	for k, v := range tc {
		// algo func
		out := twoSum60(v.in)

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
func twoSum(n int) []float64 {
	// 20200424
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.2 MB,击败了100.00% 的Go用户
	count := make(map[int]int, 0) // 统计每个点数出现的次数
	// 模拟掷第一个骰子
	for i := 1; i <= 6; i++ { // 1-6每个点数都可能出现
		count[i] = 1
	}
	// 模拟掷后边的n-1个骰子
	for i := 1; i < n; i++ {
		nextCount := map[int]int{}
		for k, v := range count {
			for i := 1; i <= 6; i++ {
				nextCount[k+i] += v
			}
		}
		count = nextCount
	}
	sum := int(math.Pow(float64(6), float64(n))) // 或者遍历count累加其值也可得到
	result := make([]float64, n*5+1)
	for k, v := range count {
		result[k-n] = float64(v) / float64(sum)
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
//
//
//
// 你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。
//
//
//
// 示例 1:
//
// 输入: 1
//输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]
//
//
// 示例 2:
//
// 输入: 2
//输出: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0
//.05556,0.02778]
//
//
//
// 限制：
//
// 1 <= n <= 11
//
