package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

type IO struct {
	in  string
	out string
}

func TestAlgo(t *testing.T) {
	// add test cases for yourself
	tc := map[string]IO{
		"1": IO{"abbbc", "a3c"},
		"2": IO{"aabbc", "aa2c"},
		"3": IO{"ab", "ab"},
		"4": IO{"abc", "abc"},
		"5": IO{"aaabc", "aaabc"},
		"6": IO{"abaabc", "a4c"},
	}

	for k, v := range tc {
		fmt.Printf("case-%s\n", k)
		out := Algo(v.in)
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

// 保留收尾，压缩中间字符，如果首字符重复，则重复输出，如果压缩后字符串没少，则保留原字符串
func Algo(in string) (out string) {
	c := in[0]
	out += string(c)
	head := true
	count := 0
	for i := 1; i < len(in); i++ {
		if i == len(in)-1 {
			if count != 0 {
				out += strconv.Itoa(count)
			}
			out += string(in[i])
			//return
		}
		//fmt.Printf("\tint[%d]=%s\n", i, string(in[i]))
		if in[i] == c && head == true {
			out += string(c)
		} else {
			head = false
			//if count == 0 && len(in) - i <= 2 {
			//	for j := i; j < len(in);j++ {
			//		out += string(in[j])
			//	}
			//	return
			//}
		}
		if !head {
			count++
		}
	}

	fmt.Printf("\tin=%s\tout=%s\n", in, out)
	if len(out) >= len(in) {
		return in
	}
	return out
}
