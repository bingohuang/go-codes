package main

import (
	"reflect"
	"testing"
)

type IO struct {
	in1 []int
	in2 int
	out []int
}

func TestTwoSum1(t *testing.T) {
	// add test cases for yourself
	tc := map[string]IO{
		"1": IO{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for k, v := range tc {
		out := twoSum1(v.in1, v.in2)
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}
