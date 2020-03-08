# Go 实现(思路+贪心算法)

> [376. 摆动序列](https://leetcode-cn.com/problems/wiggle-subsequence/)
> [完整代码实现](https://github.com/bingohuang/go-codes/blob/master/leetcode/editor/cn/p376_WiggleSubsequence_test.go)

## 思路
### 贪心规律
当序列有一段连续的递增(或递减)时，为形成摇摆子序列，我们只需要 保留这段连续的递增(或递减)的首尾元素，这样更可能使得尾部的后一个元素成为摇摆子序列的下一个元素。

如下图所示:

![](https://leetcode-cn.oss-cn-hangzhou.aliyuncs.com/p371/p371-1.png)

### 状态机

![](https://leetcode-cn.oss-cn-hangzhou.aliyuncs.com/p371/p371-2.png)


## 1. 贪心算法
```go
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	begin, up, down := 0, 1, 2
	state := begin
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		switch state {
		case begin:
			if nums[i] > nums[i-1] {
				state = up
				maxLen++
			} else if nums[i] < nums[i-1] {
				state = down
				maxLen++
			}
		case up:
			if nums[i] < nums[i-1] {
				state = down
				maxLen++
			}
		case down:
			if nums[i] > nums[i-1] {
				state = up
				maxLen++
			}
		}
	}

	return maxLen
}
```
### - 复杂度分析
- 时间复杂度：O(N)
- 空间复杂度：O(1)