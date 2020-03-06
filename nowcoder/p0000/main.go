package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// xxx.
// https://nowcoder.com/xxx
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

		MainAlgo()
	}

}

//
func MainAlgo() int {

	return 0
}
