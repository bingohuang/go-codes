# [Leetcode editor](https://plugins.jetbrains.com/plugin/12132-leetcode-editor/) customer code template

## Code File Name
```text
p$!{question.frontendQuestionId}_d$!{question.level}_$!velocityTool.camelCaseName(${question.titleSlug})_test
```

## Code Template
```text
// github.com/bingohuang/go-codes
/**
题目: ${question.frontendQuestionId}. ${question.title}
难度: $!{question.level}
地址: https://leetcode-cn.com/problems/$!{question.titleSlug}/
日期：$!velocityTool.date()
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO$!{question.frontendQuestionId} struct {
	in []string
	out []int
}

func Test$!{question.frontendQuestionId}(t *testing.T) {
	// add test cases
	tc := map[string]IO$!{question.frontendQuestionId}{
		"0": {[]string{}, []int{}},
	}

	for k, v := range tc {
		// algo func
		out := p$!{question.frontendQuestionId}(v.in)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput: %v\n", v.in)
		fmt.Printf("\toutput: %v\n", out)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

${question.code}

/* 题目详情 */
${question.content}
```
参考变量:
```text
${question.title}	question title	ex:Two Sum
${question.titleSlug}	question title slug 	ex:two-sum
${question.frontendQuestionId}	question serial number
${question.content}	question content
${question.code}	question code
$!velocityTool.camelCaseName(str)	transform str camel case
```

## live template
```
date
// $date$ @bingohuang
// 算法：$1$
// 复杂度：$2$
// 效率：
$3$
date("yyyy-MM-dd HH:mm")
```