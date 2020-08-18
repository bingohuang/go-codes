// github.com/bingohuang/go-codes
/**
题目: 54. 螺旋矩阵
难度: 2
地址: https://leetcode-cn.com/problems/spiral-matrix/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO54 struct {
	in  [][]int
	out []int
}

func Test54(t *testing.T) {
	// add test cases
	tc := map[string]IO54{
		//"0": {[][]int{}, []int{}},
		//"1": {[][]int{
		//	{ 1, 2, 3 },
		//	{ 4, 5, 6 },
		//	{ 7, 8, 9 },
		//}, []int{1,2,3,6,9,8,7,4,5}},
		//"2": {[][]int{
		//	 {1, 2, 3, 4},
		//	 {5, 6, 7, 8},
		//	 {9,10,11,12},
		//}, []int{1,2,3,4,8,12,11,10,9,5,6,7}},
		"3*": {[][]int{{3}, {2}}, []int{3, 2}},
	}

	for k, v := range tc {
		// algo func
		out := spiralOrder(v.in)

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
func spiralOrder(matrix [][]int) []int {
	// 20200818
	// 1、按层遍历
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2 MB,击败了50.00% 的Go用户
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns            = len(matrix), len(matrix[0])
		res                      = make([]int, rows*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)
	for left <= right && top <= bottom {
		for col := left; col <= right; col++ {
			res[index] = matrix[top][col]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			res[index] = matrix[row][right]
			index++
		}
		if right > left && bottom > top {
			for col := right - 1; col > left; col-- {
				res[index] = matrix[bottom][col]
				index++
			}
			for row := bottom; row > top; row-- {
				res[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
//
// 示例 1:
//
// 输入:
//[
// [ 1, 2, 3 ],
// [ 4, 5, 6 ],
// [ 7, 8, 9 ]
//]
//输出: [1,2,3,6,9,8,7,4,5]
//
//
// 示例 2:
//
// 输入:
//[
//  [1, 2, 3, 4],
//  [5, 6, 7, 8],
//  [9,10,11,12]
//]
//输出: [1,2,3,4,8,12,11,10,9,5,6,7]
//
// Related Topics 数组
// 👍 461 👎 0
