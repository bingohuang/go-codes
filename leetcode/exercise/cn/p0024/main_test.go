package main

import (
	"fmt"
	"reflect"
	"testing"
)

type IO struct {
	in  *ListNode
	out *ListNode
}

// test template
func TestAlgo(t *testing.T) {
	// add test cases for yourself
	l1I := &ListNode{Val: 1}
	l1I.Next = &ListNode{Val: 2}
	l1I.Next.Next = &ListNode{Val: 3}
	l1I.Next.Next.Next = &ListNode{Val: 4}

	l1O := &ListNode{Val: 2}
	l1O.Next = &ListNode{Val: 1}
	l1O.Next.Next = &ListNode{Val: 4}
	l1O.Next.Next.Next = &ListNode{Val: 3}

	tc := map[string]IO{
		"1": IO{l1I, l1O},
	}

	for k, v := range tc {
		out := swapPairs(v.in)
		fmt.Println("out:", out.Val, out.Next.Val, out.Next.Next.Val, out.Next.Next.Next.Val)
		fmt.Println("v.out:", v.out.Val, v.out.Next.Val, v.out.Next.Next.Val, v.out.Next.Next.Next.Val)
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, *v.out, *out)
		}
	}
}
