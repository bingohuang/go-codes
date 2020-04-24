// github.com/bingohuang/go-codes
/**
题目: 640. 求解方程
难度: 2
地址: https://leetcode-cn.com/problems/solve-the-equation/
*/
package test

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// input and ouput
type IO640 struct {
	in  string
	out string
}

func Test640(t *testing.T) {
	// add test cases
	tc := map[string]IO640{
		"1": {"x+5-3+x=6+x-2", "x=2"},
		"2": {"x=x", "Infinite solutions"},
		"3": {"2x=x", "x=0"},
		"4": {"2x+3x-6x=x+2", "x=-1"},
		"5": {"x=x+2", "No solution"},
		"6": {"0x=9", "No solution"},
	}

	for k, v := range tc {
		// algo func
		out := solveEquation(v.in)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput: %v\n", v.in)
		fmt.Printf("\toutput: %v\n", out)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func solveEquation(equation string) string {
	// 20200424
	// 执行耗时:0 ms,击败了100.00% 的Go用户
	// 内存消耗:1.9 MB,击败了100.00% 的Go用户
	// 分别对方程两侧统计x的系数和及数字的和，处理成ax+b=cx+d的形式
	exps := strings.Split(equation, "=")
	a, b := help(exps[0])
	c, d := help(exps[1])
	if a-c == 0 && d-b == 0 {
		return "Infinite solutions"
	}
	if a-c == 0 {
		return "No solution"
	}
	return fmt.Sprintf("x=%d", (d-b)/(a-c))
}

// § 预处理表达式
// 给所有"-“前边加一个”+"，然后用"+"切分整个表达式，方便处理
func help(exp string) (int, int) {
	// 给所有"-"前边加一个"+"，方便处理
	exp = strings.ReplaceAll(exp, "-", "+-")
	s := strings.Split(exp, "+")
	xCount, num := 0, 0
	for _, v := range s {
		if len(v) > 0 && v[len(v)-1] == 'x' {
			if v == "x" {
				xCount++
			} else if v == "-x" {
				xCount--
			} else {
				n, _ := strconv.Atoi(v[:len(v)-1])
				xCount += n
			}
		} else {
			n, _ := strconv.Atoi(v)
			num += n
		}
	}
	return xCount, num
}

//使用正则切分表达式
//与给上一个方法思路一致，只是不再修改原表达式，直接用正则切分；只有前两行代码不一样：
func help1(exp string) (int, int) {
	re, _ := regexp.Compile("[+-]?[0-9]*x?")
	s := re.FindAllString(exp, -1)
	xCount, num := 0, 0
	for _, v := range s {
		if len(v) > 0 && v[len(v)-1] == 'x' {
			if v == "x" || v == "+x" {
				xCount++
			} else if v == "-x" {
				xCount--
			} else {
				n, _ := strconv.Atoi(v[:len(v)-1])
				xCount += n
			}
		} else {
			n, _ := strconv.Atoi(v)
			num += n
		}
	}
	return xCount, num
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//求解一个给定的方程，将x以字符串"x=#value"的形式返回。该方程仅包含'+'，' - '操作，变量 x 和其对应系数。
//
// 如果方程没有解，请返回“No solution”。
//
// 如果方程有无限解，则返回“Infinite solutions”。
//
// 如果方程中只有一个解，要保证返回值 x 是一个整数。
//
// 示例 1：
//
// 输入: "x+5-3+x=6+x-2"
//输出: "x=2"
//
//
// 示例 2:
//
// 输入: "x=x"
//输出: "Infinite solutions"
//
//
// 示例 3:
//
// 输入: "2x=x"
//输出: "x=0"
//
//
// 示例 4:
//
// 输入: "2x+3x-6x=x+2"
//输出: "x=-1"
//
//
// 示例 5:
//
// 输入: "x=x+2"
//输出: "No solution"
//
// Related Topics 数学
