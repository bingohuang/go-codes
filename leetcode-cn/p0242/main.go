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
