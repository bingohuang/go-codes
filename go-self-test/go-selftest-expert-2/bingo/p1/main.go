package main

import "fmt"

/*
	题目大意：
		有一个长度为N的数组[]nums，数组内第i个数的大小是nums[i]，现在你要对每个数求出 在他左边离他最近的 比他小的数。

	数据范围：1<= N, nums[i] <= 100000

	样例1：
		输入：[1,6,4,10,2,5]
		输出：[-1,1,1,4,1,2]
*/
func main() {
	case1 := []int{1, 6, 4, 10, 2, 5}
	fmt.Println("case1:", case1, "\n\tright:", []int{-1, 1, 1, 4, 1, 2}, "\n\twrong:", leftMinNum(case1))

}

func leftMinNum(T []int) []int {
	var r []int
	var i, j int
	for i = 0; i < len(T); i++ {
		for j = i - 1; j >= 0; j-- {
			if T[j] < T[i] {
				r = append(r, T[j])
				break
			}
		}
		if j == -1 {
			r = append(r, -1)
		}
	}
	return r
}
