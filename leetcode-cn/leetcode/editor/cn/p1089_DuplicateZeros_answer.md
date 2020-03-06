# Go 语言的三种实现

> [1089. 复写零](https://leetcode-cn.com/problems/duplicate-zeros/)

## 1. 暴力破解法
```go
// 开辟一个新的slice，复制
func duplicateZeros(arr []int) {
	t := make([]int, len(arr))
	i := 0
	for _, v := range arr {
		if i == len(arr) {
			break
		}
		if v != 0 {
			t[i] = v
			i++
		} else {
			t[i] = 0
			i++
			if i == len(arr) {
				break
			}
			t[i] = 0
			i++
		}
	}
	fmt.Printf("t=%v\narr=%v\n", t, arr)
	copy(arr, t)
}
```

## 2. 两次遍历法
```go
// 算法2：两次遍历法
func duplicateZeros(arr []int) {
	// 1. 计算需要复写0的数量
	pDups := 0
	length := len(arr) - 1
	for left := 0; left <= length-pDups; left++ {
		if arr[left] == 0 {
			// 2. 注意处理元素边界上0的情况
			if left == length-pDups {
				arr[length] = 0
				length--
				break
			}
			pDups++
		}
	}

	// 3. 从末尾迭代数组
	last := length - pDups
	for i := last; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+pDups] = 0
			pDups--
			arr[i+pDups] = 0
		} else {
			arr[i+pDups] = arr[i]
		}
	}
}
```

## 3. append复制法(借用Slice特性，只需8行代码)
```go
// 算法3： append复制转移法
func duplicateZeros(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			i++
		}
	}
}
```