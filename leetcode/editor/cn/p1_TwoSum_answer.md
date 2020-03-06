# Go 两种实现

> https://leetcode-cn.com/problems/two-sum/solution/go-liang-chong-shi-xian-by-bingohuang/

> [1. 两数之和](https://leetcode-cn.com/problems/two-sum/solution/)

## 1. 暴力破解法
```go
// 算法1：暴力破解法
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
```
## 2. map法
// 算法2：map法
```go
func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}
```