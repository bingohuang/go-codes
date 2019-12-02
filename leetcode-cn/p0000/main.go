package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// xxxx. template
// https://leetcode-cn.com/problems/xxxx
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

		fmt.Println(algo(n, m))
	}

}

func algo(n int, m int) int {

	return n + m
}
