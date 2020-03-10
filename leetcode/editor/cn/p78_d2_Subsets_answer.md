# Go 实现(回溯)

> [78. 子集(中等)](https://leetcode-cn.com/problems/subsets/)

> [完整代码实现](https://github.com/bingohuang/go-codes/blob/master/leetcode/editor/cn/p78_d2_Subsets_test.go)

## 思路 1: 回溯法
利用回溯法生成子集，即对于每个元素，都有试探放入或不放入集合中的两个选择，这个过程称为回溯试探法

### Go 实现1：回溯, 两次递归
```go
// 算法1：回溯, 两次递归，注意传递指针，否则会导致slice扩容，结果不对
func subsets1(nums []int) [][]int {
	result := make([][]int, 0)
	item := make([]int, 0)

	result = append(result, item)
	generate1(0, nums, &item, &result)
	return result
}

func generate1(i int, nums []int, item *[]int, result *[][]int) {
	if i >= len(nums) {
		return
	}
	*item = append(*item, nums[i])
	temp := make([]int, len(*item))
	for i, v := range *item {
		temp[i] = v
	}
	*result = append(*result, temp)
	generate1(i+1, nums, item, result)
	*item = (*item)[:len(*item)-1]
	generate1(i+1, nums, item, result)
	return
}
```

### Go 实现1：回溯，一层递归+一套循环
```go
// 算法2：回溯，一层递归+一套循环，此处使用闭包函数
func subsets2(nums []int) [][]int {
	result := make([][]int, 0)

	var generate2 func(pos int, num int, item []int)
	generate2 = func(pos int, num int, item []int) {
		if len(item) == num {
			tmp := make([]int, len(item))
			copy(tmp, item)
			result = append(result, tmp)
			return
		}
		for i := pos; i < len(nums); i++ { // 注意：小于nums长度
			item = append(item, nums[i])
			generate2(i+1, num, item)
			item = item[:len(item)-1]
		}
	}

	for i := 0; i <= len(nums); i++ {
		item := make([]int, 0, i) // 注意cap
		generate2(0, i, item)
	}
	return result
}
```
### 复杂度分析
- 时间复杂度：$O(N×2^N)$
- 空间复杂度：$O(N×2^N)$