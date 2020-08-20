// github.com/bingohuang/go-codes
/**
题目: 122. 买卖股票的最佳时机 II
难度: 1
地址: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
日期：2020-08-20 10:44:59
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO122 struct {
	in  []int
	out int
}

func Test122(t *testing.T) {
	// add test cases
	tc := map[string]IO122{
		//"0": {[]int{}, 0},
		"1": {[]int{7, 1, 5, 3, 6, 4}, 7},
		"2": {[]int{1, 2, 3, 4, 5}, 4},
		"3": {[]int{7, 6, 4, 3, 1}, 0},
	}

	for k, v := range tc {
		// algo func
		out := maxProfit(v.in)

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
func maxProfit(prices []int) int {
	// 2020-08-20 10:54 @bingohuang
	// 算法：1、峰谷法
	// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-ii-by-leetcode/
	// 复杂度：O(N)/O(1)
	// 效率：执行耗时:4 ms,击败了95.45% 的Go用户
	//		内存消耗:3.1 MB,击败了100.00% 的Go用户
	max := 0
	n := len(prices)
	valley := prices[0]
	peak := prices[0]
	for i := 0; i < n; i++ {
		// 找谷值
		for ; i < n-1 && prices[i] >= prices[i+1]; i++ {
		}
		valley = prices[i]
		// 找峰值
		for ; i < n-1 && prices[i] <= prices[i+1]; i++ {
		}
		peak = prices[i]
		max += peak - valley
	}
	// 上面11行可以等价为如下代码
	/*for i := 0; i < n-1; i++ {
		if prices[i] < prices[i+1] {
			max += prices[i+1] - prices[i]
		}
	}*/
	return max
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//
//
// 示例 1:
//
// 输入: [7,1,5,3,6,4]
//输出: 7
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
//
//
// 示例 2:
//
// 输入: [1,2,3,4,5]
//输出: 4
//解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
//     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
//
//
// 示例 3:
//
// 输入: [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
//
//
//
// 提示：
//
//
// 1 <= prices.length <= 3 * 10 ^ 4
// 0 <= prices[i] <= 10 ^ 4
//
// Related Topics 贪心算法 数组
// 👍 821 👎 0
