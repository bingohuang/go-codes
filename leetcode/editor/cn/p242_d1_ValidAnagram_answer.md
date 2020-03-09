# Go 两种实现

> [242. 有效的字母异位词 - 简单](https://leetcode-cn.com/problems/valid-anagram/)

> [完整代码实现](https://github.com/bingohuang/go-codes/blob/master/leetcode/editor/cn/p242_d1_ValidAnagram_test.go)

## 1. 两个数组比较法
```go
// 算法1：
// 可以利用两个长度都是26的字符串数组，
// 统计每个字符串中小写字母出现的次数，
// 然后再对比是否相等
func isAnagram1(s string, t string) bool {
	a := [26]int{}
	b := [26]int{}

	for _, v := range s {
		a[v-'a'] += 1
	}

	for _, v := range t {
		b[v-'a'] += 1
	}

	fmt.Printf("a=%v\nb=%v\n", a, b)

	return a == b
}
```

## 2. 一个数组判零法
// 算法2：map法
```go
// 算法2：
// 可以只利用一个长度为 26 的字符数组，
// 将出现在字符串 s 里的字符个数加 1，
// 而出现在字符串 t 里的字符个数减 1，
// 最后判断每个小写字母的个数是否都为 0。
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arr := [26]int{}
	for _, v := range s {
		arr[v-'a']++
	}
	for _, v := range t {
		arr[v-'a']--
	}
	fmt.Printf("arr=%v\n", arr)
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}
```
### 复杂度分析 - 以上两种算法复杂度相同
- 时间复杂度：$O(N)$ 因为访问计数器表是一个固定的时间操作。
- 空间复杂度：$O(1)$ 尽管我们使用了额外的空间，但是空间的复杂性是 $O(1)$，因为无论 $N$ 有多大，表的大小都保持不变。