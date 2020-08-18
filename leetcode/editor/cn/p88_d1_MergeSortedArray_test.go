// github.com/bingohuang/go-codes
/**
题目: 88. 合并两个有序数组
难度: 1
地址: https://leetcode-cn.com/problems/merge-sorted-array/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO88 struct {
	in1 []int
	in2 int
	in3 []int
	in4 int
	out []int
}

func Test88(t *testing.T) {
	// add test cases
	//nums1 = [1,2,3,0,0,0], m = 3
	//nums2 = [2,5,6],       n = 3
	//输出: [1,2,2,3,5,6]
	tc := map[string]IO88{
		"1":  {[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		"2*": {[]int{2, 0}, 1, []int{1}, 1, []int{1, 2}},
	}

	for k, v := range tc {
		// algo func
		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput1: %v\n", v.in1)
		fmt.Printf("\tinput2: %v\n", v.in2)
		fmt.Printf("\tinput3: %v\n", v.in3)
		fmt.Printf("\tinput4: %v\n", v.in4)
		merge(v.in1, v.in2, v.in3, v.in4)
		fmt.Printf("\toutput: %v\n", v.in1)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(v.in1, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, v.in1)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 20200818
	// 1、合并再排序
	// O((n+m)log(n+m))/O(1)
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.3 MB,击败了25.33% 的Go用户
	/*nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)*/
	// 2、双指针：从前往后
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.6 MB,击败了5.00% 的Go用户
	/*var stack []int

	p1, p2 := 0,0
	for p1 < m && p2 < n {
		if nums1[p1] < nums2[p2] {
			stack = append(stack, nums1[p1])
			p1 ++
		} else {
			stack = append(stack, nums2[p2])
			p2 ++
		}
	}
	if p1 < m {
		stack = append(stack, nums1[p1:m]...)
	}
	if p2 < n {
		stack = append(stack, nums2[p2:n]...)
	}
	for k, v := range stack {
		nums1[k] = v
	}
	fmt.Println(nums1)*/
	// 3、双指针，从后往前
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.3 MB,击败了67.67% 的Go用户
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		// 这么干不行，传出去的值不对
		//nums1 = append(nums2[:n], nums1[n:]...)
		//fmt.Println(nums1)
	}
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
//
//
//
// 说明:
//
//
// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
//
//
//
//
// 示例:
//
// 输入:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6]
// Related Topics 数组 双指针
// 👍 594 👎 0
