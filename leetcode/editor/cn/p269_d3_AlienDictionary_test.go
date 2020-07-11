// github.com/bingohuang/go-codes
/**
题目: 269. 火星词典
难度: 3
地址: https://leetcode-cn.com/problems/alien-dictionary/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO269 struct {
	in  []string
	out string
}

func Test269(t *testing.T) {
	// add test cases
	tc := map[string]IO269{
		"1": {[]string{}, ""},
	}

	for k, v := range tc {
		// algo func
		out := alienOrder(v.in)

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
func alienOrder(words []string) string {

	return ""
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//现有一种使用字母的全新语言，这门语言的字母顺序与英语顺序不同。
//
// 假设，您并不知道其中字母之间的先后顺序。但是，会收到词典中获得一个 不为空的 单词列表。因为是从词典中获得的，所以该单词列表内的单词已经 按这门新语言的字
//母顺序进行了排序。
//
// 您需要根据这个输入的列表，还原出此语言中已知的字母顺序。
//
//
//
// 示例 1：
//
// 输入:
//[
//  "wrt",
//  "wrf",
//  "er",
//  "ett",
//  "rftt"
//]
//输出: "wertf"
//
//
// 示例 2：
//
// 输入:
//[
//  "z",
//  "x"
//]
//输出: "zx"
//
//
// 示例 3：
//
// 输入:
//[
//  "z",
//  "x",
//  "z"
//]
//输出: ""
//解释: 此顺序是非法的，因此返回 ""。
//
//
//
//
// 提示：
//
//
// 你可以默认输入的全部都是小写字母
// 若给定的顺序是不合法的，则返回空字符串即可
// 若存在多种可能的合法字母顺序，请返回其中任意一种顺序即可
//
// Related Topics 图 拓扑排序
