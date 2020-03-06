package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// xxx.
// https://leetcode-cn.com/problems/xxx
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

//给你一堆抢票时间，求抢到票的人数
//每秒都有人抢到票
//如果一秒内时间最早的一个人算抢到
//同一秒最早同一时间有多人算多个
//类似
//2019-08-11 12:12:11.001
//2019-08-11 12:12:11.135
//2019-08-11 12:12:11.653
//2019-08-11 12:13:12.001
//2019-08-11 12:13:12.001
//2019-08-11 12:13:12.002
//2019-08-11 12:13:13.002
//有3个人抢到
func Shakedown(times []string) int {
	var count int
	timeMap := make(map[string]string)
	for _, time := range times {
		s := strings.Split(time, ".")
		fmt.Println("s:", s)
		fmt.Println("s[0]:", s[0])
		fmt.Println("s[1]:", s[1])
		if v, ok := timeMap[s[0]]; !ok {
			fmt.Println("v:", v)
			fmt.Println("ok :", ok)
			timeMap[s[0]] = s[1]
			fmt.Println("timeMap1:", timeMap)
			count++
			fmt.Println("count1:", count)
		} else if ok && v == s[1] {
			fmt.Println("timeMap2:", timeMap)
			count++
			fmt.Println("count2:", count)
		}
	}
	return count
}
