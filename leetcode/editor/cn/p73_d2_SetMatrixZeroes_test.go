// github.com/bingohuang/go-codes
/**
题目: 73. 矩阵置零
难度: 2
地址: https://leetcode-cn.com/problems/set-matrix-zeroes/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO73 struct {
	in  [][]int
	out [][]int
}

func Test73(t *testing.T) {
	// add test cases
	// 输入:
	//[
	//  [1,1,1],
	//  [1,0,1],
	//  [1,1,1]
	//]
	//输出:
	//[
	//  [1,0,1],
	//  [0,0,0],
	//  [1,0,1]
	//]
	// 输入:
	//[
	//  [0,1,2,0],
	//  [3,4,5,2],
	//  [1,3,1,5]
	//]
	//输出:
	//[
	//  [0,0,0,0],
	//  [0,4,5,0],
	//  [0,3,1,0]
	//]
	tc := map[string]IO73{
		//"0": {[][]int{}, [][]int{}},
		"1": {[][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}, [][]int{
			{1, 0, 1},
			{0, 0, 0},
			{1, 0, 1},
		}},
		"2": {[][]int{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		}, [][]int{
			{0, 0, 0, 0},
			{0, 4, 5, 0},
			{0, 3, 1, 0},
		}},
	}

	for k, v := range tc {
		// algo func

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput: %v\n", v.in)
		setZeroes(v.in)
		fmt.Printf("\toutput: %v\n", v.in)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(v.in, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, v.in)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func setZeroes(matrix [][]int) {
	// 20200819
	// 1、全部复制
	// O(M*N)
	// 不需要用visitor数组，直接在原数组设置
	/*if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	m := len(matrix)
	n := len(matrix[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if !visited[i][j] && matrix[i][j] == 0 {
				for p := 0; p<n;p++ {
					if !visited[i][p] {
						matrix[i][p] = 0
						visited[i][p] = true
					}
				}
				for p := 0; p<m;p++ {
					if !visited[p][j] {
						matrix[p][j] = 0
						visited[p][j] = true
					}
				}
			}
			//visited[i][j] = true
		}
	}
	fmt.Println(matrix)*/
	// https://leetcode-cn.com/problems/set-matrix-zeroes/solution/ju-zhen-zhi-ling-by-leetcode/
	// 1、两次扫描，额外空间，用set
	// O(M*N)/O(M+N)
	// 执行耗时:16 ms,击败了70.23% 的Go用户
	// 内存消耗:6 MB,击败了42.53% 的Go用户
	// 本题算法不需要，工程化最好加上
	//if len(matrix) == 0 || len(matrix[0]) == 0 {
	//	return
	//}
	/*m := len(matrix)
	n := len(matrix[0])
	rows := map[int]bool{}
	cols := map[int]bool{}
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for i:=0;i<m;i++ {
		for j := 0; j < n; j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}*/
	// 3、O(1)空间复杂度+暴力
	// O((MxN)x(M+N))/O(1)
	// 执行耗时:24 ms,击败了7.36% 的Go用户
	// 内存消耗:6 MB,击败了13.79% 的Go用户
	/*FLAG := math.MinInt64
	m := len(matrix)
	n := len(matrix[0])
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if matrix[i][j] == 0{
				for k :=0;k<m;k++{
					if matrix[k][j] != 0{
						matrix[k][j] = FLAG
					}
				}
				for k :=0;k<n;k++{
					if matrix[i][k] != 0 {
						matrix[i][k] = FLAG
					}
				}
			}
		}
	}

	for i:=0;i<m;i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == FLAG {
				matrix[i][j] = 0
			}
		}
	}*/
	// 3、只将每行和每列的第一个元素赋值为0
	// O(MxN)/O(1)
	// 执行耗时:20 ms,击败了15.72% 的Go用户
	// 内存消耗:6.1 MB,击败了5.75% 的Go用户
	m := len(matrix)
	n := len(matrix[0])
	isCol := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			isCol = true
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if matrix[0][0] == 0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if isCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	fmt.Println(matrix)
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。
//
// 示例 1:
//
// 输入:
//[
//  [1,1,1],
//  [1,0,1],
//  [1,1,1]
//]
//输出:
//[
//  [1,0,1],
//  [0,0,0],
//  [1,0,1]
//]
//
//
// 示例 2:
//
// 输入:
//[
//  [0,1,2,0],
//  [3,4,5,2],
//  [1,3,1,5]
//]
//输出:
//[
//  [0,0,0,0],
//  [0,4,5,0],
//  [0,3,1,0]
//]
//
// 进阶:
//
//
// 一个直接的解决方案是使用 O(mn) 的额外空间，但这并不是一个好的解决方案。
// 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
// 你能想出一个常数空间的解决方案吗？
//
// Related Topics 数组
// 👍 274 👎 0
