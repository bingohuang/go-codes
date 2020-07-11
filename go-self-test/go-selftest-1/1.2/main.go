package main

import (
	"fmt"
	"sort"
)

func main() {
	var input []int
	var max int
	for {
		var n int
		if _, err := fmt.Scanf("%d", &n); err != nil {
			break
		}
		input = append(input, n)
		if max < n {
			max = n
		}
	}
	n := len(input) - 1
	limitedHours := input[n]
	trees := input[:n]
	fmt.Println(leastK(trees, max, limitedHours))
}

// 朴素实现
// 假设桃子最多的树上桃子数为max，显然答案在[1,max]区间里，从1开始向max递增尝试即可
// 时间复杂度O(n*max) n为桃树棵数，max为桃子最多那棵树上的桃子数；空间复杂度O(1)
func leastK(trees []int, max, limitedHours int) int {
	n := len(trees)
	if n > limitedHours {
		return -1
	}
	if n == limitedHours {
		return max
	}
	for k := 1; k < max; k++ {
		if costHoursForK(trees, k) <= limitedHours {
			return k
		}
	}
	return -1
}

// 二分搜索
// k从1到max去遍历，太老实了，如 2 3 800 5 这样的输入，有太多多余的尝试
// 这里可以用二分查找的思路，将复杂度变成O(nlog(max))
func leastK2(trees []int, max, H int) int {
	n := len(trees)
	if H < n {
		return -1
	}
	if H == n {
		return max
	}
	left, right := 1, max+1
	for left < right {
		mid := left + (right-left)/2
		if costHoursForK(trees, mid) <= H { // 可能有多个值使得costHoursForK(trees, mid) == H， 需要最左侧的那个值
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 还可以使用标准库sort.Search，减少代码量，同时体会下将比较逻辑抽离的思想：
func leastK3(trees []int, max, H int) int {
	n := len(trees)
	if H < n {
		return -1
	}
	if H == n {
		return max
	}
	return sort.Search(max, func(i int) bool { // 注意取值为max的情况已经在前边返回，所以不再包含max
		if i == 0 {
			return false
		}
		return costHoursForK(trees, i) <= H
	})
}

func costHoursForK(trees []int, k int) int {
	sum := 0
	for _, v := range trees {
		sum += (v-1)/k + 1
		/* 另一个实现
		   sum += v / k
		   if v%k > 0 {
		   	sum += 1
		   }
		*/
	}
	return sum
}
