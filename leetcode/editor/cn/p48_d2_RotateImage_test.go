// github.com/bingohuang/go-codes
/**
题目: 48. 旋转图像
难度: 2
地址: https://leetcode-cn.com/problems/rotate-image/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO48 struct {
	in  [][]int
	out [][]int
}

func Test48(t *testing.T) {
	// add test cases
	tc := map[string]IO48{
		"0": {[][]int{}, [][]int{}},
		"1": {[][]int{{1}}, [][]int{{1}}},
		"2": {[][]int{{1, 2}, {3, 4}}, [][]int{{3, 1}, {4, 2}}},
		"3": {
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9}},
			[][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3}}},
	}

	for k, v := range tc {
		// algo func

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput: %v\n", v.in)
		rotate(v.in)
		fmt.Printf("\toutput: %v\n", v.in)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(v.in, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, v.in)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func rotate(matrix [][]int) {
	// 20200818
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.2 MB,击败了100.00% 的Go用户
	n := len(matrix)
	// transpose matrix
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// reverse each row
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个 n × n 的二维矩阵表示一个图像。
//
// 将图像顺时针旋转 90 度。
//
// 说明：
//
// 你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
//
// 示例 1:
//
// 给定 matrix =
//[
//  [1,2,3],
//  [4,5,6],
//  [7,8,9]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//  [7,4,1],
//  [8,5,2],
//  [9,6,3]
//]
//
//
// 示例 2:
//
// 给定 matrix =
//[
//  [ 5, 1, 9,11],
//  [ 2, 4, 8,10],
//  [13, 3, 6, 7],
//  [15,14,12,16]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//  [15,13, 2, 5],
//  [14, 3, 4, 1],
//  [12, 6, 8, 9],
//  [16, 7,10,11]
//]
//
// Related Topics 数组
// 👍 534 👎 0
