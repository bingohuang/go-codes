// github.com/bingohuang/go-codes
/**
题目: 4. 寻找两个有序数组的中位数
难度: 3
地址: https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO4 struct {
	in1 []int
	in2 []int
	out float64
}

func Test4(t *testing.T) {
	// add test cases
	tc := map[string]IO4{
		"1": {[]int{1, 3}, []int{2}, 2},
		"2": {[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for k, v := range tc {
		// algo func
		out := findMedianSortedArrays(v.in1, v.in2)

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
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 20200423
	// 解法一：暴力法，归并合并，求中位数
	// 执行耗时:20 ms,击败了54.61% 的Go用户
	// 内存消耗:6 MB,击败了18.18% 的Go用户
	// 执行耗时:16 ms,击败了79.17% 的Go用户
	// 内存消耗:6 MB,击败了18.18% 的Go用户
	res := merge4(nums1, nums2)
	if len(res)%2 == 0 {
		return float64(res[len(res)/2-1]+res[len(res)/2]) / 2.0
	} else {
		return float64(res[len(res)/2])
	}

	// 解法二：切分法
	// 解法三：数组组合

}

func merge4(left []int, right []int) []int {
	l, r := 0, 0
	var res []int
	// 注意：[左右]对比，是指左的第一个元素，与右边的第一个元素进行对比，
	// 哪个小，就先放到结果的第一位，然后左或右取出了元素的那边的索引进行++
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			res = append(res, left[l])
			//因为处理了左边的第r个元素，所以r的指针要向前移动一个单位
			l++
		} else {
			res = append(res, right[r])
			//因为处理了右边的第r个元素，所以r的指针要向前移动一个单位
			r++
		}
	}
	// 比较完后，还要分别将左，右的剩余的元素，追加到结果列的后面。
	res = append(res, left[l:]...)
	res = append(res, right[r:]...)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//
// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
// 你可以假设 nums1 和 nums2 不会同时为空。
//
// 示例 1:
//
// nums1 = [1, 3]
//nums2 = [2]
//
//则中位数是 2.0
//
//
// 示例 2:
//
// nums1 = [1, 2]
//nums2 = [3, 4]
//
//则中位数是 (2 + 3)/2 = 2.5
//
// Related Topics 数组 二分查找 分治算法
