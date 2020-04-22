// github.com/bingohuang/go-codes
/**
题目: 215. 数组中的第K个最大元素
难度: 2
地址: https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
*/
package test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// input and ouput
type IO215 struct {
	in1 []int
	in2 int
	out int
}

func Test215(t *testing.T) {
	// add test cases
	tc := map[string]IO215{
		"1": {[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		"2": {[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}

	for k, v := range tc {
		// algo func
		out := findKthLargest(v.in1, v.in2)

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
func findKthLargest(nums []int, k int) int {
	// 20200422: 解法 1：直接将数组进行排序，然后得出结果。
	// 执行耗时:8 ms,击败了92.06% 的Go用户
	// 内存消耗:3.5 MB,击败了100.00% 的Go用户
	sort.Ints(nums)
	return nums[len(nums)-k]
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 示例 1:
//
// 输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//
//
// 示例 2:
//
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
//
// 说明:
//
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
// Related Topics 堆 分治算法
