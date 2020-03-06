package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 猴子吃桃

//比如，输入：1 2 3 4 5 6
//期望输出：4
//五颗桃树，桃子个数分别为1，2，3，4，5，天神6小时来，猴子最慢吃桃速度为4。
func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		input, _ := r.ReadString('\n')
		if input == "" || input == "\n" {
			return
		}
		input = strings.TrimSpace(input)
		// Added
		//如果只有一个输入，就没有桃子树的数量或者时间，无解
		if len(input) <= 1 {
			fmt.Println(-1)
			continue
		}

		inputs := strings.Split(input, " ")
		//fmt.Println("inputs:", inputs)

		var piles []int // 堆数
		var reasonable bool
		for i := 0; i < len(inputs)-1; i++ {
			p, _ := strconv.Atoi(inputs[i])
			// Added
			// 如果有树的堆数小于零，不合理，这条记录输出-1
			if p < 0 {
				reasonable = false
				break
			}
			piles = append(piles, p)
		}
		if !reasonable {
			fmt.Println(-1)
			continue
		}
		h, _ := strconv.Atoi(inputs[len(inputs)-1]) // 时间

		// Added
		//时间小于桃子树的颗数，桃子是吃不完的，无解
		if h < len(piles) {
			fmt.Println(-1)
			continue
		}

		//fmt.Println("piles:", piles)
		//fmt.Println("h:", h)
		fmt.Println(minEatingSpeed(piles, h))

	LOOP:
	}

}

//有蟠桃树若干，每棵树上有N个蟠桃，N > 0，天神把守。
//猴子喜欢慢慢吃桃，每小时吃K个蟠桃，而且这一小时只在一棵树上吃。
//如果此树桃不足K，猴子吃完，剩下时间在树上玩耍，不再吃。
//天神H小时候到来。
//已知每棵桃树桃子个数，求在天神到来前猴子吃完桃子最小速度K。
func minEatingSpeed(piles []int, h int) int {
	lo := 1
	hi := maxPile(piles)

	for lo < hi {
		mi := (lo + hi) / 2
		if canEat(piles, h, mi) {
			hi = mi
		} else {
			lo = mi + 1
		}
	}

	return lo
}

// 求最大的蟠桃数
func maxPile(piles []int) int {
	max := 0
	for _, p := range piles {
		if p > max {
			max = p
		}
	}
	return max
}

// 计算是否可能在规定的时间(h)以指定的速度(k)内吃完所有蟠桃(piles)
func canEat(piles []int, h int, k int) bool {
	time := 0
	for _, p := range piles {
		time += (p-1)/k + 1
	}
	return time <= h
}
