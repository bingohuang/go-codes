# Go 实现

> [860. 柠檬水找零 - 简单](https://leetcode-cn.com/problems/lemonade-change/)

> [完整代码实现](https://github.com/bingohuang/go-codes/blob/master/leetcode/editor/cn/p860_d1_LemonadeChange_test.go)

## 1. 遍历法
```go
// 算法1：
// 如果给5元，不用找
// 如果给10元，找5元，否则找不开
// 如果给20元，优先找10+5，再找5+5+5，否则找不开
func lemonadeChange1(bills []int) bool {
	five, ten := 0, 0

	for _, v := range bills {
		if v == 5 {
			five ++
		} else if v == 10 {
			if five >= 1 {
				five --
				ten ++
			} else {
				return false
			}
		} else if v == 20 {
			if five >= 1 && ten >= 1 {
				five --
				ten --
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}

	return true
}
```
### 复杂度分析
- 时间复杂度：$O(N)$
- 空间复杂度：$O(1)$