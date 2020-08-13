// github.com/bingohuang/go-codes
/**
题目: 15. 三数之和
难度: 2
地址: https://leetcode-cn.com/problems/3sum/
*/
package test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// input and ouput
type IO15 struct {
	in  []int
	out [][]int
}

func Test15(t *testing.T) {
	// add test cases
	// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
	//满足要求的三元组集合为：
	//  [-1, 0, 1],
	//  [-1, -1, 2]
	tc := map[string]IO15{
		"0": {[]int{}, [][]int{}},
		"1": {[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, 0, 1}, {-1, 2, -1}}},
	}

	for k, v := range tc {
		// algo func
		out := threeSum(v.in)

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
func threeSum(nums []int) [][]int {
	// 20200813
	// 1、暴力法：要注意去重，可能还得用hash表来去重，得不偿失，不可取
	// O(N^2)/O(logN)
	// 执行耗时:40 ms,击败了43.98% 的Go用户
	// 内存消耗:6.9 MB,击败了59.33% 的Go用户
	var res [][]int
	// 先排序
	sort.Ints(nums)
	// Time Limit Exceeded
	/*for i:=0; i < len(nums); i ++ {
		if i == 0 || nums[i] != nums[i-1] {
			for j := i + 1;j < len(nums);j ++ {
				if j == i+1 || nums[j] != nums[j-1] {
					for k := j + 1; k < len(nums); k++ {
						if (k == j + 1 || nums[k] != nums[k-1]) && nums[i]+nums[j]+nums[k] == 0 {
							res = append(res, []int{nums[i], nums[j], nums[k]})
						}
					}
				}
			}

		}
	}*/
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := n - 1
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[i]+nums[j]+nums[k] > 0 {
				k--
			}
			if j == k {
				break
			}
			if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复
//的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例：
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
//
// Related Topics 数组 双指针
// 👍 2462 👎 0
