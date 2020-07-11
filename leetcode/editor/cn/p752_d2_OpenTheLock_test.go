// github.com/bingohuang/go-codes
/**
题目: 752. 打开转盘锁
难度: 2
地址: https://leetcode-cn.com/problems/open-the-lock/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO752 struct {
	in1 []string
	in2 string
	out int
}

func Test752(t *testing.T) {
	// add test cases
	tc := map[string]IO752{
		//"1": IO752{[]string{"0201","0101","0102","1212","2002"}, "0202", 6},
		"2": IO752{[]string{"8888"}, "0009", 1},
	}

	for k, v := range tc {
		// algo func
		out := openLock(v.in1, v.in2)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput1: %v\n", v.in1)
		fmt.Printf("\tinput2: %v\n", v.in2)
		fmt.Printf("\toutput: %v\n", out)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func openLock(deadends []string, target string) int {
	step := 0
	// 记录需要跳过的死亡密码
	deads := make(map[string]struct{})
	for _, s := range deadends {
		deads[s] = struct{}{}
	}
	if _, ok := deads["0000"]; ok {
		return -1
	}
	if target == "0000" {
		return 0
	}

	// 记录已经穷举过的密码，防止走回头路
	visited := make(map[string]struct{})
	queue := make([]string, 0)
	queue = append(queue, "0000")
	visited["0000"] = struct{}{}
	var l int
	for len(queue) > 0 {
		// 增加步数
		step++
		l = len(queue)
		for i := 0; i < l; i++ {
			// 弹出
			cur := queue[0]

			// 判断是否是死亡密码
			if _, ok := deads[cur]; ok {
				continue
			}
			// 判断是否达到终点
			if cur == target {
				return step
			}

			// 将一个节点的未遍历相邻节点加入队列
			for j := 0; j < 4; j++ {
				// 上拨
				up := plusOne(cur, j)
				if _, ok := visited[up]; !ok {
					queue = append(queue, up)
					visited[up] = struct{}{}
				}
				// 下拨
				down := minusOne(cur, j)
				if _, ok := visited[down]; !ok {
					queue = append(queue, down)
					visited[down] = struct{}{}
				}
			}
		}
		queue = queue[l:]
	}

	return -1
}
func plusOne(s string, j int) string {
	res := []byte(s)
	if res[j] == '9' {
		res[j] = '0'
	} else {
		res[j]++
	}
	return string(res)
}
func minusOne(s string, j int) string {
	res := []byte(s)
	if res[j] == '0' {
		res[j] = '9'
	} else {
		res[j]--
	}
	return string(res)
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'
// 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
//
// 锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
//
// 列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
//
// 字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。
//
//
//
// 示例 1:
//
//
//输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
//输出：6
//解释：
//可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
//注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
//因为当拨动到 "0102" 时这个锁就会被锁定。
//
//
// 示例 2:
//
//
//输入: deadends = ["8888"], target = "0009"
//输出：1
//解释：
//把最后一位反向旋转一次即可 "0000" -> "0009"。
//
//
// 示例 3:
//
//
//输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], targ
//et = "8888"
//输出：-1
//解释：
//无法旋转到目标数字且不被锁定。
//
//
// 示例 4:
//
//
//输入: deadends = ["0000"], target = "8888"
//输出：-1
//
//
//
//
// 提示：
//
//
// 死亡列表 deadends 的长度范围为 [1, 500]。
// 目标数字 target 不会在 deadends 之中。
// 每个 deadends 和 target 中的字符串的数字会在 10,000 个可能的情况 '0000' 到 '9999' 中产生。
//
// Related Topics 广度优先搜索
