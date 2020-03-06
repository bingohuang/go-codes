package main

import (
	"reflect"
	"testing"
)

type IO struct {
	in1 int
	in2 int
	out int
}

// test template
func TestAlgo(t *testing.T) {
	// add test cases for yourself
	tc := map[string]IO{
		"1": IO{1, 2, 3},
		"2": IO{2, 3, 5},
		"3": IO{3, 4, 7},
	}

	for k, v := range tc {
		out := algo(v.in1, v.in2)
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}
