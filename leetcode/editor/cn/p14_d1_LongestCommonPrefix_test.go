// github.com/bingohuang/go-codes
/**
题目: 14. 最长公共前缀
难度: 1
地址: https://leetcode-cn.com/problems/longest-common-prefix/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO14 struct {
	in  []string
	out string
}

func Test14(t *testing.T) {
	// add test cases
	tc := map[string]IO14{
		"0": {[]string{}, ""},
		"1": {[]string{"flower", "flow", "flight"}, "fl"},
		"2": {[]string{"dog", "racecar", "car"}, ""},
		"3": {[]string{"hu", "he", "hi"}, "h"},
	}

	for k, v := range tc {
		// algo func
		out := longestCommonPrefix(v.in)

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
func longestCommonPrefix(strs []string) string {
	// 20200810
	// 1、暴力法
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.4 MB,击败了56.62% 的Go用户
	/*res := ""
	if len(strs) < 1 {
		return res
	}
	str0 := strs[0]
	get := false
	for i:=len(str0);i > 0 && !get ;i--{
		res = str0[:i]
		get = true
		for j:=1; j < len(strs); j ++ {
			if !strings.HasPrefix(strs[j], res) {
				get = false
				break
			}
		}
	}
	if get {
		return res
	}
	return ""*/

	// 2、公式法：LCP(S1​…Sn​)=LCP(LCP(LCP(S1​,S2​),S3​),…Sn​)
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.4 MB,击败了88.60% 的Go用户
	/*if len(strs) == 0 {
		return ""
	}
	lcp := func(str1, str2 string) string {
		length := len(str1)
		if len(str1) > len(str2) {
			length = len(str2)
		}
		index := 0
		for index < length && str1[index] == str2[index] {
			index ++
		}
		return str1[:index]
	}
	prefix := strs[0]
	for i:=1;i<len(strs);i++{
		prefix=lcp(prefix, strs[i])
		if len(prefix)==0{
			return ""
		}
	}
	return prefix*/

	// 3、纵向扫描法，也是第一个方法的一个代码改进，更简洁
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:2.4 MB,击败了88.60% 的Go用户
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]

}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
// 示例 1:
//
// 输入: ["flower","flow","flight"]
//输出: "fl"
//
//
// 示例 2:
//
// 输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//
//
// 说明:
//
// 所有输入只包含小写字母 a-z 。
// Related Topics 字符串
// 👍 1200 👎 0
