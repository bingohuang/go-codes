package main

import "testing"

type IO struct {
	records []Record
	balance int
}

func TestName(t *testing.T) {
	// add test cases for yourself
	tc := map[string]IO{
		"1": IO{[]Record{
			{120, 10, 12},
		}, 238},
		"2": IO{[]Record{
			{120, 10, 12},
			{120, 25, 25},
		}, 133},
	}

	for k, v := range tc {
		out := calcBalance(v.records)
		if out != v.balance {
			t.Errorf("case-%v: except answer: [%v, %v], get answer: [%v, %v]", k, v.records, v.balance, v.records, out)
		}
	}
}
