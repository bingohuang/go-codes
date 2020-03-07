# [Leetcode editor](https://plugins.jetbrains.com/plugin/12132-leetcode-editor/) customer code template

## Code File Name
```text
p$!{question.frontendQuestionId}_$!velocityTool.camelCaseName(${question.titleSlug})_test
```

## Code Template
```text
// github.com/bingohuang/go-codes
/**
题目: $!{question.formTitle}
难度: $!{question.level}
地址: https://leetcode-cn.com/problems/$!{question.titleSlug}/
*/
package test

import (
	"fmt"
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

		fmt.Println("input :", v.in)
		fmt.Println("output:", out)
		fmt.Println("except:", v.out)

		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

${question.code}

/* 题目详情 */
${question.content}
```