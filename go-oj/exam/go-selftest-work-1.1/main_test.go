package main

import "testing"

type IO struct {
	in1 []int
	in2 int
	out int
}

func TestName(t *testing.T) {
	// add test cases for yourself
	tc := map[string]IO{
		"1": IO{[]int{3, 6, 7, 11}, 8, 4},
		"2": IO{[]int{30, 11, 23, 4, 20}, 5, 30},
		"3": IO{[]int{30, 11, 23, 4, 20}, 6, 23},
	}

	for k, v := range tc {
		out := minEatingSpeed(v.in1, v.in2)
		if out != v.out {
			t.Errorf("case-%v: except answer: [%v, %v,  %v], get answer: [%v, %v,  %v]", k, v.in1, v.in1, v.out, v.in1, v.in1, out)
		}
	}
}
