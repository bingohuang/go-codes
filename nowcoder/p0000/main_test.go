package main

import (
	"reflect"
	"testing"
)

type IO struct {
	in  []string
	out int
}

func TestAlgo(t *testing.T) {
	// add test cases for yourself
	tc := map[string]IO{
		"1": IO{[]string{}, 0},
	}

	for k, v := range tc {
		out := Algo()
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

//
func Algo() int {

	return 0
}
