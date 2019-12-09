package main

import (
	"fmt"
	"strings"
)

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
//结果：有4个人抢到
//给定的抢票时间是按照先后顺序，通过map的方式判断是否有最早时刻有一起抢到票
func Shakedown(times []string) int {
	var count int
	timeMap := make(map[string]string)
	for _, time := range times {
		s := strings.Split(time, ".")
		fmt.Println("s:", s)
		if v, ok := timeMap[s[0]]; !ok {
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
	fmt.Println("count:", count)
	return count
}
