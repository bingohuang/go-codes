package main

import "fmt"

func allocSlice(min, high int) []int {
	var b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 99: 100}
	fmt.Printf("slice b: len(%d), cap(%d)\n",
		len(b), cap(b))

	return b[min:high] // 使得额外的96个整型占用的空间无法及时释放

	//nb := make([]int, high-min, high-min)
	//copy(nb, b[min:high]) // 还可以避免"隐匿"数据暴露
	//return nb
}

func main() {
	b1 := allocSlice(3, 7)
	fmt.Printf("slice b1: len(%d), cap(%d), elements(%v)\n",
		len(b1), cap(b1), b1)
}
