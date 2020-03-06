/*
	题目大意：
		你有一个长度为N的数组[]nums，数组内第i个数的大小是nums[i]，现在你要对每个数求出 在他左边离他最近的 比他小的数。

	数据范围：1<= N, nums[i] <= 100000

	样例1：
		输入：[1,6,4,10,2,5]
		输出：[-1,1,1,4,1,2]
*/

package main

import (
	"fmt"
)

func leftNum(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	stack := make([]int, n) // 维护一个递增的单调栈
	ans := make([]int, n)
	top := 0
	for i := 0; i < n; i++ {
		for top > 0 && nums[stack[top-1]] >= nums[i] { // 当前元素和栈顶比较，弹出不满足单调性的栈顶
			top--
		}
		if top == 0 {
			ans[i] = -1
		} else {
			ans[i] = nums[stack[top-1]]
		}
		stack[top] = i // 当前元素入栈
		top++
	}
	return ans
}

func main() {
	nums := []int{1, 6, 4, 10, 2, 5}
	ans := leftNum(nums)
	fmt.Println(ans)
}
