package main

import (
	"reflect"
	"testing"
	"time"
)

type IO struct {
	in1 string
	in2 string
	out bool
}

// test template
func Test(t *testing.T) {
	// add test cases for yourself
	tc := map[string]IO{
		"1": IO{"anagram", "nagaram", true},
		"2": IO{"rat", "car", false},
		"3": IO{"bingo", "bing", false},
		"4": IO{"bingohuang", "huangbingo", true},
	}

	for k, v := range tc {
		out := isAnagram(v.in1, v.in2)
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		} else {
			t.Logf("case-%v: OK", k)
		}
	}
	time.Sleep(1)
}
