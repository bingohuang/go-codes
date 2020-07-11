// github.com/bingohuang/go-codes
/**
题目: 499. 迷宫 III
难度: 3
地址: https://leetcode-cn.com/problems/the-maze-iii/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO499 struct {
	in1 [][]int
	in2 []int
	in3 []int
	out string
}

func Test499(t *testing.T) {
	// add test cases
	tc := map[string]IO499{
		"1": {[][]int{}, []int{}, []int{}, ""},
	}

	for k, v := range tc {
		// algo func
		out := findShortestWay(v.in1, v.in2, v.in3)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput1: %v\n", v.in1)
		fmt.Printf("\tinput2: %v\n", v.in2)
		fmt.Printf("\tinput3: %v\n", v.in3)
		fmt.Printf("\toutput: %v\n", out)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func findShortestWay(maze [][]int, ball []int, hole []int) string {
	// 20200424
	// 首先参考迷宫II，但重点注意：洞不不需要靠墙，切要保证字典顺序
	return ""
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//由空地和墙组成的迷宫中有一个球。球可以向上（u）下（d）左（l）右（r）四个方向滚动，但在遇到墙壁前不会停止滚动。当球停下时，可以选择下一个方向。迷宫中还有
//一个洞，当球运动经过洞时，就会掉进洞里。
//
// 给定球的起始位置，目的地和迷宫，找出让球以最短距离掉进洞里的路径。 距离的定义是球从起始位置（不包括）到目的地（包括）经过的空地个数。通过'u', 'd'
//, 'l' 和 'r'输出球的移动方向。 由于可能有多条最短路径， 请输出字典序最小的路径。如果球无法进入洞，输出"impossible"。
//
// 迷宫由一个0和1的二维数组表示。 1表示墙壁，0表示空地。你可以假定迷宫的边缘都是墙壁。起始位置和目的地的坐标通过行号和列号给出。
//
//
//
// 示例1:
//
// 输入 1: 迷宫由以下二维数组表示
//
//0 0 0 0 0
//1 1 0 0 1
//0 0 0 0 0
//0 1 0 0 1
//0 1 0 0 0
//
//输入 2: 球的初始位置 (rowBall, colBall) = (4, 3)
//输入 3: 洞的位置 (rowHole, colHole) = (0, 1)
//
//输出: "lul"
//
//解析: 有两条让球进洞的最短路径。
//第一条路径是 左 -> 上 -> 左, 记为 "lul".
//第二条路径是 上 -> 左, 记为 'ul'.
//两条路径都具有最短距离6, 但'l' < 'u'，故第一条路径字典序更小。因此输出"lul"。
//
//
//
// 示例 2:
//
// 输入 1: 迷宫由以下二维数组表示
//
//0 0 0 0 0
//1 1 0 0 1
//0 0 0 0 0
//0 1 0 0 1
//0 1 0 0 0
//
//输入 2: 球的初始位置 (rowBall, colBall) = (4, 3)
//输入 3: 洞的位置 (rowHole, colHole) = (3, 0)
//
//输出: "impossible"
//
//示例: 球无法到达洞。
//
//
//
//
//
// 注意:
//
//
// 迷宫中只有一个球和一个目的地。
// 球和洞都在空地上，且初始时它们不在同一位置。
// 给定的迷宫不包括边界 (如图中的红色矩形), 但你可以假设迷宫的边缘都是墙壁。
// 迷宫至少包括2块空地，行数和列数均不超过30。
//
// Related Topics 深度优先搜索 广度优先搜索
