// github.com/bingohuang/go-codes
/**
题目: 剑指 Offer 20. 表示数值的字符串
难度: 2
地址: https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/
日期: 2020-09-02 21:21:49
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO剑指Offer20 struct {
	in  string
	out bool
}

func Test剑指Offer20(t *testing.T) {
	// add test cases
	tc := map[string]IO剑指Offer20{
		"0": IO剑指Offer20{"0", true},
		"1": IO剑指Offer20{"+100", true},
		"2": IO剑指Offer20{"5e2", true},
		"3": IO剑指Offer20{"12e", false},
		"4": IO剑指Offer20{"1a3.14", false},
	}

	for k, v := range tc {
		// algo func
		out := isNumber(v.in)

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
func isNumber(s string) bool {
	// 2020-09-02 21:49 @bingohuang
	// 算法：1、状态机法
	// 复杂度：O(N)/O(1)
	// 效率：Runtime:4 ms, faster than 61.48% of Go online submissions.
	//			Memory Usage:5.7 MB, less than 11.86% of Go online submissions.
	type state int
	type char int
	const (
		state_initial state = iota
		state_int_sign
		state_integer
		state_point
		state_point_without_int
		state_fraction
		state_exp
		state_exp_sign
		state_exp_number
		state_end
	)
	const (
		char_number char = iota
		char_exp
		char_point
		char_sign
		char_space
		char_illegal
	)

	toChar := func(ch byte) char {
		switch ch {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return char_number
		case 'e', 'E':
			return char_exp
		case '.':
			return char_point
		case '+', '-':
			return char_sign
		case ' ':
			return char_space
		default:
			return char_illegal
		}
	}

	transfer := map[state]map[char]state{
		state_initial: map[char]state{
			char_space:  state_initial,
			char_number: state_integer,
			char_point:  state_point_without_int,
			char_sign:   state_int_sign,
		},
		state_int_sign: map[char]state{
			char_number: state_integer,
			char_point:  state_point_without_int,
		},
		state_integer: map[char]state{
			char_number: state_integer,
			char_exp:    state_exp,
			char_point:  state_point,
			char_space:  state_end,
		},
		state_point: map[char]state{
			char_number: state_fraction,
			char_exp:    state_exp,
			char_space:  state_end,
		},
		state_point_without_int: map[char]state{
			char_number: state_fraction,
		},
		state_fraction: map[char]state{
			char_number: state_fraction,
			char_exp:    state_exp,
			char_space:  state_end,
		},
		state_exp: map[char]state{
			char_number: state_exp_number,
			char_sign:   state_exp_sign,
		},
		state_exp_sign: map[char]state{
			char_number: state_exp_number,
		},
		state_exp_number: map[char]state{
			char_number: state_exp_number,
			char_space:  state_end,
		},
		state_end: map[char]state{
			char_space: state_end,
		},
	}
	myState := state_initial
	for i := 0; i < len(s); i++ {
		typ := toChar(s[i])
		if _, ok := transfer[myState][typ]; !ok {
			return false
		} else {
			myState = transfer[myState][typ]
		}
	}
	return myState == state_integer || myState == state_point ||
		myState == state_fraction || myState == state_exp_number || myState == state_end
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"-1E-16"、"012
//3"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。
//
//
// Related Topics 数学
// 👍 86 👎 0
