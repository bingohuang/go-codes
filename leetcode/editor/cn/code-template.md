# code template

## code file name
```text
p$!{question.frontendQuestionId}_$!velocityTool.camelCaseName(${question.titleSlug})_test
```

## code content
```text
// github.com/bingohuang/go-codes
package test

import (
	"reflect"
	"testing"
)

// input and ouput
type IO$!{question.frontendQuestionId} struct {
	in []int
	out []int
}

func Test$!{question.frontendQuestionId}(t *testing.T) {
	// add test cases
	tc := map[string]IO$!{question.frontendQuestionId}{
		"1": IO$!{question.frontendQuestionId}{[]int{}, []int{}},
	}

	for k, v := range tc {
		// algo func
		out := p$!{question.frontendQuestionId}(v.in)
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

${question.code}

${question.content}

```