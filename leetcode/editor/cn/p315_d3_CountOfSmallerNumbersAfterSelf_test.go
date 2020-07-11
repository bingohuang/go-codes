// github.com/bingohuang/go-codes
/**
题目: 315. 计算右侧小于当前元素的个数
难度: 3
地址: https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO315 struct {
	in  []int
	out []int
}

func Test315(t *testing.T) {
	// add test cases
	tc := map[string]IO315{
		"1": IO315{[]int{5, 2, 6, 1}, []int{2, 1, 1, 0}},
		"2": IO315{[]int{6, 6, 6, 1, 1, 1}, []int{3, 3, 3, 0, 0, 0}},
		"3": IO315{[]int{5, -7, 9, 1, 3, 5, -2, 1}, []int{5, 0, 5, 1, 2, 2, 0, 0}},
	}

	for k, v := range tc {
		// algo func
		out := countSmaller(v.in)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput: %v\n", v.in)
		fmt.Printf("\toutput: %v\n", out)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func countSmaller(nums []int) []int {

	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是 nums[i] 右侧小于 num
//s[i] 的元素的数量。
//
// 示例:
//
// 输入: [5,2,6,1]
//输出: [2,1,1,0]
//解释:
//5 的右侧有 2 个更小的元素 (2 和 1).
//2 的右侧仅有 1 个更小的元素 (1).
//6 的右侧有 1 个更小的元素 (1).
//1 的右侧有 0 个更小的元素.
//
// Related Topics 排序 树状数组 线段树 二分查找 分治算法
