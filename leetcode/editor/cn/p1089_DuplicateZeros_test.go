// github.com/bingohuang/go-codes
/**
题目: [1089]复写零
难度: 1
地址: https://leetcode-cn.com/problems/duplicate-zeros/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO1089 struct {
	in  []int
	out []int
}

func Test1089(t *testing.T) {
	// add test cases
	tc := map[string]IO1089{
		"1": IO1089{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		"2": IO1089{[]int{1, 2, 3}, []int{1, 2, 3}},
		"3": IO1089{[]int{1, 2, 0}, []int{1, 2, 0}},
		"4": IO1089{[]int{8, 4, 5, 0, 0, 0, 0, 7}, []int{8, 4, 5, 0, 0, 0, 0, 0}},
	}

	for k, v := range tc {
		// algo func
		duplicateZeros3(v.in)
		if !reflect.DeepEqual(v.in, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, v.in)
		}
	}
}

// 算法1
// 开辟一个新的slice，复制
func duplicateZeros1(arr []int) {
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

// 算法2：两次遍历法
func duplicateZeros2(arr []int) {
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

// 算法3： append复制转移法
func duplicateZeros3(arr []int) {
	fmt.Printf("arr=%v\n", arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			fmt.Printf("\tarr=%v\n", arr)
			i++
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func duplicateZeros(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			i++
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给你一个长度固定的整数数组 arr，请你将该数组中出现的每个零都复写一遍，并将其余的元素向右平移。
//
// 注意：请不要在超过该数组长度的位置写入元素。
//
// 要求：请对输入的数组 就地 进行上述修改，不要从函数返回任何东西。
//
//
//
// 示例 1：
//
// 输入：[1,0,2,3,0,4,5,0]
//输出：null
//解释：调用函数后，输入的数组将被修改为：[1,0,0,2,3,0,0,4]
//
//
// 示例 2：
//
// 输入：[1,2,3]
//输出：null
//解释：调用函数后，输入的数组将被修改为：[1,2,3]
//
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 10000
// 0 <= arr[i] <= 9
//
// Related Topics 数组
