// github.com/bingohuang/go-codes
/**
题目: 118. 杨辉三角
难度: 1
地址: https://leetcode-cn.com/problems/pascals-triangle/
日期：2020-08-20 09:51:20
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO118 struct {
	in  int
	out [][]int
}

func Test118(t *testing.T) {
	// add test cases
	tc := map[string]IO118{
		"0": {0, nil},
		"1": {5, [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		}},
	}

	for k, v := range tc {
		// algo func
		out := generate(v.in)

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
func generate(numRows int) [][]int {
	// 2020-08-20 10:01 @bingohuang
	// 算法：1、找规律
	// 复杂度：O(numRows^2)/O(numRows^2)
	// 效率：执行耗时:0 ms,击败了100.00% 的Go用户
	//		内存消耗:2 MB,击败了89.35% 的Go用户
	var ans [][]int
	if numRows <= 0 {
		// 以下两种返回，leetcode都AC
		//return [][]int{}
		//return nil
		return ans
	}
	ans = append(ans, []int{1})
	for i := 1; i < numRows; i++ {
		list := []int{1}
		// 优化本段代码
		/*up := len(ans[i-1])
		if up >= 2 {
			for j := 0; j < up-1; j ++ {
				list = append(list, ans[i-1][j] + ans[i-1][j+1])
			}
		}*/
		for j := 1; j < i; j++ {
			list = append(list, ans[i-1][j-1]+ans[i-1][j])
		}
		list = append(list, 1)
		ans = append(ans, list)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
//
//
//
// 在杨辉三角中，每个数是它左上方和右上方的数的和。
//
// 示例:
//
// 输入: 5
//输出:
//[
//     [1],
//    [1,1],
//   [1,2,1],
//  [1,3,3,1],
// [1,4,6,4,1]
//]
// Related Topics 数组
// 👍 339 👎 0
