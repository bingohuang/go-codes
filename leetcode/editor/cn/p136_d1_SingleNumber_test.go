// github.com/bingohuang/go-codes
/**
题目: 136. 只出现一次的数字
难度: 1
地址: https://leetcode-cn.com/problems/single-number/
日期：2020-08-20 19:01:27
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO136 struct {
	in  []int
	out int
}

func Test136(t *testing.T) {
	// add test cases
	tc := map[string]IO136{
		"0": {[]int{0}, 0},
		"1": {[]int{2, 2, 1}, 1},
		"2": {[]int{4, 1, 2, 1, 2}, 4},
	}

	for k, v := range tc {
		// algo func
		out := singleNumber(v.in)

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
func singleNumber(nums []int) int {
	// 2020-08-20 19:10 @bingohuang
	// 算法：1、map+循环
	// 复杂度：O(N)/O(N)
	// 效率：执行耗时:16 ms,击败了19.64% 的Go用户
	//			内存消耗:5.9 MB,击败了8.50% 的Go用户
	/*m := make(map[int]int)
	for i := 0;i<len(nums);i++{
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = 1
		} else {
			//m[nums[i]] ++
			// 可以用删除的方法: 可是也并没有省空间呢
			// 执行耗时:16 ms,击败了19.64% 的Go用户
			// 内存消耗:5.9 MB,击败了9.95% 的Go用户
			delete(m, nums[i])
		}
	}
	//for k, v := range m {
	//	if v == 1 {
	//		return k
	//	}
	//
	// 删除法
	for k := range m {
			return k
	}
	return 0*/
	// 2020-08-20 19:23 @bingohuang
	// 算法：2、异或法，可以得到O(1)的空间复杂度
	// 复杂度：O(N)/O(1)
	// 效率：执行耗时:12 ms,击败了75.97% 的Go用户
	//			内存消耗:4.7 MB,击败了100.00% 的Go用户
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
// 说明：
//
// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
// 示例 1:
//
// 输入: [2,2,1]
//输出: 1
//
//
// 示例 2:
//
// 输入: [4,1,2,1,2]
//输出: 4
// Related Topics 位运算 哈希表
// 👍 1441 👎 0
