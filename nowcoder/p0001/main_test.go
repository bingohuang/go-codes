package main

import (
	"reflect"
	"testing"
)

type IO struct {
	in  []string
	out int
}

func TestShakedown(t *testing.T) {
	// add test cases for yourself
	tc := map[string]IO{
		"1": {[]string{
			"2019-08-11 12:12:11.001",
			"2019-08-11 12:12:11.135",
			"2019-08-11 12:12:11.653",
			"2019-08-11 12:13:12.001",
			"2019-08-11 12:13:12.001",
			"2019-08-11 12:13:12.002",
			"2019-08-11 12:13:13.002",
		}, 4},
	}
	for k, v := range tc {
		out := Shakedown(v.in)
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}
