// github.com/bingohuang/go-codes
/**
题目: 121. 买卖股票的最佳时机
难度: 1
地址: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO121 struct {
	in  []int
	out int
}

func Test121(t *testing.T) {
	// add test cases
	tc := map[string]IO121{
		"1": IO121{[]int{}, 0},
	}

	for k, v := range tc {
		// algo func
		out := maxProfit1(v.in)

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
// TODO: 提交前记得改名
func maxProfit1(prices []int) int {

	return 0
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
// 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
//
// 注意：你不能在买入股票前卖出股票。
//
//
//
// 示例 1:
//
// 输入: [7,1,5,3,6,4]
//输出: 5
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
//
//
// 示例 2:
//
// 输入: [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
//
// Related Topics 数组 动态规划