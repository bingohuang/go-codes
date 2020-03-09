# Go 实现(栈+贪心)

> [402. 移掉K位数字 - 中等](https://leetcode-cn.com/problems/remove-k-digits/)

> [完整代码实现](https://github.com/bingohuang/go-codes/blob/master/leetcode/editor/cn/p402_d2_RemoveKDigits_test.go)

## 思路 1: 栈的贪心算法

### 分析
若去掉某一位数字，为了使得到的新数字最小，需要**尽可能**让得到的新数字**优先最高位**最小，**其次次高位**最小，**再其次第3位**最小...**依次循环**

### 贪心规律
从**高位向低位**遍历，如果对应的数字**大于**下一位数字，则把该位数字去掉，得到的数字**最小**。

去掉 K 个数字，需要从最高位遍历 k 次

### 注意
1. 当所有数字都扫描完成后， k 仍然大于 0，该如何处理？ 如: num=12345, k=3
- 从尾部删除剩余 k 个数字

2. 当数字中有 0 出现，该如何特殊处理？如: num=100200， k = 1
- 如果保留的数字串非空，则保留，否则不保留（比如 00200 无意义）。详情见代码注释

3. 如何将最后结果存储为字符串并返回？
- Go 中，可以使用 `string('0' + number)` 的方式转化为字符

### Go 实现
```go
func removeKdigits(num string, k int) string {
	var stack []uint8 // 使用 slice 作为栈
	var result string // 存储最终结果字符串
	for i := 0; i < len(num); i++ { // 从最高位扫描数字字符串num
		number := num[i] - '0' // 将字符串转化为整数
		// 条件：当栈不为空，且栈顶元素大于下一个数字，且 k 还大于0
		for len(stack) != 0 && stack[len(stack)-1] > number && k > 0 {
			stack = stack[:len(stack)-1] // 弹出栈顶元素
			k-- // 剔除一个数字，k 减一
		}
		// 条件：如果数字不为空，或者数字为空但是栈不为空
		if number != 0 || len(stack) != 0 {
			stack = append(stack, number) // 入栈
		}
	}
	// 条件：如果栈不空且还能删除数字
	for len(stack) != 0 && k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	// 将栈中元素转化为字符串
	for _, v := range stack {
		result += string('0' + v)
	}
	// 注意为空情况
	if result == "" {
		return "0"
	}
	return result
}
```
### 复杂度分析
- 时间复杂度：$O(N*K)$ 其中 $0<k≤N$
- 空间复杂度：$O(N)$ 在最坏的情况下栈存储了所有的数字