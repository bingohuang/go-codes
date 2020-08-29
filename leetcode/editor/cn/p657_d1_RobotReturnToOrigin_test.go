// github.com/bingohuang/go-codes
/**
题目: 657. 机器人能否返回原点
难度: 1
地址: https://leetcode-cn.com/problems/robot-return-to-origin/
日期：2020-08-28 08:48:50
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO657 struct {
	in  string
	out bool
}

func Test657(t *testing.T) {
	// add test cases
	// R（右），L（左），U（上）和 D（下）
	tc := map[string]IO657{
		//"0": {"", true},
		//"1": {"UD", true},
		"2": {"LL", false},
	}

	for k, v := range tc {
		// algo func
		out := judgeCircle(v.in)

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
func judgeCircle(moves string) bool {
	// 2020-08-28 08:50 @bingohuang
	// 算法：1、遍历法，判断左和右，上和下是否数量一样
	// 复杂度：O(N)/O(N)
	// 效率：执行耗时:0 ms,击败了100.00% 的Go用户
	//		内存消耗:3.5 MB,击败了43.48% 的Go用户
	// R（右），L（左），U（上）和 D（下）
	/*a := [4]int{}
	for _, s := range moves {
		switch s {
		case 'R':
			a[0] ++
		case 'L':
			a[1] ++
		case 'U':
			a[2] ++
		case 'D':
			a[3] ++
		}
	}
	return a[0] == a[1] && a[2] == a[3]*/

	// 2020-08-29 11:24 @bingohuang
	// 算法： 模拟自己人自己行走，看最后是否会走到远点
	// 复杂度：O(N)/O(1)
	// 效率：执行耗时:4 ms,击败了92.31% 的Go用户
	//			内存消耗:3.5 MB,击败了100.00% 的Go用户
	x, y := 0, 0
	for _, s := range moves {
		switch s {
		case 'R':
			x++
		case 'L':
			x--
		case 'U':
			y++
		case 'D':
			y--
		}
	}
	return x == 0 && y == 0
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//在二维平面上，有一个机器人从原点 (0, 0) 开始。给出它的移动顺序，判断这个机器人在完成移动后是否在 (0, 0) 处结束。
//
// 移动顺序由字符串表示。字符 move[i] 表示其第 i 次移动。机器人的有效动作有 R（右），L（左），U（上）和 D（下）。如果机器人在完成所有动作后
//返回原点，则返回 true。否则，返回 false。
//
// 注意：机器人“面朝”的方向无关紧要。 “R” 将始终使机器人向右移动一次，“L” 将始终向左移动等。此外，假设每次移动机器人的移动幅度相同。
//
//
//
// 示例 1:
//
// 输入: "UD"
//输出: true
//解释：机器人向上移动一次，然后向下移动一次。所有动作都具有相同的幅度，因此它最终回到它开始的原点。因此，我们返回 true。
//
// 示例 2:
//
// 输入: "LL"
//输出: false
//解释：机器人向左移动两次。它最终位于原点的左侧，距原点有两次 “移动” 的距离。我们返回 false，因为它在移动结束时没有返回原点。
// Related Topics 字符串
// 👍 136 👎 0
