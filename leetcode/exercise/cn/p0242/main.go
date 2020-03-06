package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 242. 有效的字母异位词
// https://leetcode-cn.com/problems/valid-anagram/
func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		input, _ := r.ReadString('\n')
		if input == "" || input == "\n" {
			return
		}
		input = strings.TrimSpace(input)
		inputs := strings.Split(input, " ")
		//fmt.Println("inputs:", inputs)

		n, _ := strconv.Atoi(inputs[0])
		m, _ := strconv.Atoi(inputs[1])

		fmt.Println(n, m)
	}

}

/*
解题思路

一个重要的前提“假设两个字符串只包含小写字母”，小写字母一共也就 26 个，因此：

1. 可以利用两个长度都为 26 的字符数组来统计每个字符串中小写字母出现的次数，然后再对比是否相等；

2. 可以只利用一个长度为 26 的字符数组，将出现在字符串 s 里的字符个数加 1，而出现在字符串 t 里的字符个数减 1，最后判断每个小写字母的个数是否都为 0。

按上述操作，可得出结论：s 和 t 互为字母异位词。
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var w [26]int
	for _, i := range s {
		//fmt.Println(i)
		w[i-'a'] += 1
	}
	fmt.Println("s:", w)

	for _, i := range t {
		w[i-'a'] -= 1
	}
	fmt.Println("t:", w)

	for _, v := range w {
		if v != 0 {
			return false
		}
	}

	return true
}
