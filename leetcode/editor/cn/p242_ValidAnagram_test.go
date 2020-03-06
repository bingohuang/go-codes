package test

import (
	"fmt"
	"reflect"
	"testing"
)

type IO242 struct {
	in1 string
	in2 string
	out bool
}

func Test242(t *testing.T) {
	// add test cases for yourself
	tc := map[string]IO242{
		"1": IO242{"anagram", "nagaram", true},
		"2": IO242{"rat", "car", false},
	}

	for k, v := range tc {
		out := isAnagram2(v.in1, v.in2)
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

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

//leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
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

//leetcode submit region end(Prohibit modification and deletion)

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 示例 1:
//
// 输入: s = "anagram", t = "nagaram"
//输出: true
//
//
// 示例 2:
//
// 输入: s = "rat", t = "car"
//输出: false
//
// 说明:
//你可以假设字符串只包含小写字母。
//
// 进阶:
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
// Related Topics 排序 哈希表
