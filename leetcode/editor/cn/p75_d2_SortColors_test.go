// github.com/bingohuang/go-codes
/**
题目: 75. 颜色分类
难度: 2
地址: https://leetcode-cn.com/problems/sort-colors/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO75 struct {
	in  []int
	out []int
}

func Test75(t *testing.T) {
	// add test cases
	tc := map[string]IO75{
		"1": {[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	}

	for k, v := range tc {
		// algo func

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput: %v\n", v.in)
		sortColors(v.in)
		fmt.Printf("\toutput: %v\n", v.in)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(v.in, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, v.in)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func sortColors(nums []int) {
	// 20200424
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.1 MB,击败了100.00% 的Go用户
	red, white, blue := 0, 0, 0
	for _, v := range nums {
		if v == 0 {
			red++
		} else if v == 1 {
			white++
		} else if v == 2 {
			blue++
		}
	}
	for i := 0; i < len(nums); i++ {
		if i >= 0 && i < red {
			nums[i] = 0
		} else if i >= red && i < red+white {
			nums[i] = 1
		} else if i >= red+white && i < red+white+blue {
			nums[i] = 2
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
// 注意:
//不能使用代码库中的排序函数来解决这道题。
//
// 示例:
//
// 输入: [2,0,2,1,1,0]
//输出: [0,0,1,1,2,2]
//
// 进阶：
//
//
// 一个直观的解决方案是使用计数排序的两趟扫描算法。
// 首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
// 你能想出一个仅使用常数空间的一趟扫描算法吗？
//
// Related Topics 排序 数组 双指针
