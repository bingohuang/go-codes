// github.com/bingohuang/go-codes
/**
题目: 49. 字母异位词分组
难度: 2
地址: https://leetcode-cn.com/problems/group-anagrams/
*/
package test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// input and ouput
type IO49 struct {
	in  []string
	out [][]string
}

func Test49(t *testing.T) {
	// add test cases
	tc := map[string]IO49{
		"0": {nil, nil},
		"1": {[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				{"ate", "eat", "tea"},
				{"nat", "tan"},
				{"bat"},
			}},
	}

	for k, v := range tc {
		// algo func
		out := groupAnagrams(v.in)

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
func groupAnagrams(strs []string) [][]string {
	// 20200818
	// 1、排序法
	// O(NKlogK)/O(NK)
	// 执行耗时:32 ms,击败了48.00% 的Go用户
	// 内存消耗:8 MB,击败了35.24% 的Go用户
	var res [][]string
	mp := make(map[string][]string)
	for _, s := range strs {
		b := []byte(s)
		sort.Slice(b, func(i int, j int) bool { return b[i] < b[j] })
		ns := string(b)
		mp[ns] = append(mp[ns], s)
	}
	for _, v := range mp {
		res = append(res, v)
	}
	return res
	// 2、计数法：对每个字母做#count
	// O(NK)/)(NK)
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//
// 说明：
//
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。
//
// Related Topics 哈希表 字符串
// 👍 435 👎 0
