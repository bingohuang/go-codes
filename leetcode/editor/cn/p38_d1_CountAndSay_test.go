// github.com/bingohuang/go-codes
/**
题目: 38. 外观数列
难度: 1
地址: https://leetcode-cn.com/problems/count-and-say/
*/
package test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

// input and ouput
type IO38 struct {
	in  int
	out string
}

func Test38(t *testing.T) {
	// add test cases
	tc := map[string]IO38{
		//"0": {0, ""},
		//"1": {1, "1"},
		//"2": {2, "11"},
		"3": {3, "21"},
		//"4": {4, "1211"},
		//"5": {5, "111221"},
		//"6": {6, "312211"},
		//"7": {7, "13112221"},
	}

	for k, v := range tc {
		// algo func
		out := countAndSay(v.in)

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
func countAndSay(n int) string {
	// 20200818
	/*if n ==1 {
		return "1"
	}
	pre := "11"
	res := ""
	count := 1
	for j := 1; j <=n; j ++ {
		res = ""
		for i := 1;i<len(pre);i++{
			if pre[i] != pre[i-1] {
				res += fmt.Sprintf("%d%s", count, string(pre[i-1]))
				count = 1
			} else {
				count ++
			}
			if i+1 == len(pre) {
				res += fmt.Sprintf("%d%s", count, string(pre[i]))
				pre = res
				//return res
			}
		}
	}
	return res*/
	// 1、迭代法
	// https://leetcode-cn.com/problems/count-and-say/solution/ji-su-jie-bu-di-gui-zhi-ji-lu-qian-hou-liang-ren-p/
	// 执行耗时:4 ms,击败了48.24% 的Go用户
	// 内存消耗:6.4 MB,击败了44.97% 的Go用户
	/*prev := "1"
	for i := 1; i < n; i++ {
		next, count, num := "", 1, prev[0]
		for j := 1; j < len(prev); j++ {
			if prev[j] == num {
				count++
			} else {
				next += strconv.Itoa(count) + string(num)
				num = prev[j]
				count = 1
			}
		}
		next += strconv.Itoa(count) + string(num)
		prev = next
	}
	return prev*/
	// 2、递归法：加上 n==1判断，去掉外层循环
	// 执行耗时:4 ms,击败了48.24% 的Go用户
	// 内存消耗:6.3 MB,击败了57.14% 的Go用户
	if n == 1 {
		return "1"
	}
	prev := countAndSay(n - 1)
	next, count, num := "", 1, prev[0]
	for j := 1; j < len(prev); j++ {
		if prev[j] == num {
			count++
		} else {
			next += strconv.Itoa(count) + string(num)
			num = prev[j]
			count = 1
		}
	}
	next += strconv.Itoa(count) + string(num)
	prev = next
	return prev
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
//
// 注意：整数序列中的每一项将表示为一个字符串。
//
// 「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
//
// 1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//
//
// 第一项是数字 1
//
// 描述前一项，这个数是 1 即 “一个 1 ”，记作 11
//
// 描述前一项，这个数是 11 即 “两个 1 ” ，记作 21
//
// 描述前一项，这个数是 21 即 “一个 2 一个 1 ” ，记作 1211
//
// 描述前一项，这个数是 1211 即 “一个 1 一个 2 两个 1 ” ，记作 111221
//
//
//
// 示例 1:
//
// 输入: 1
//输出: "1"
//解释：这是一个基本样例。
//
// 示例 2:
//
// 输入: 4
//输出: "1211"
//解释：当 n = 3 时，序列是 "21"，其中我们有 "2" 和 "1" 两组，"2" 可以读作 "12"，也就是出现频次 = 1 而 值 = 2；类似
//"1" 可以读作 "11"。所以答案是 "12" 和 "11" 组合在一起，也就是 "1211"。
// Related Topics 字符串
// 👍 523 👎 0
