// github.com/bingohuang/go-codes
/**
题目: 841. 钥匙和房间
难度: 2
地址: https://leetcode-cn.com/problems/keys-and-rooms/
日期：2020-08-31 21:40:42
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO841 struct {
	in  [][]int
	out bool
}

func Test841(t *testing.T) {
	// add test cases
	tc := map[string]IO841{
		"0": {[][]int{{0}}, true},
		"1": {[][]int{{1}, {2}, {3}, {}}, true},
		"2": {[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false},
	}

	for k, v := range tc {
		// algo func
		out := canVisitAllRooms(v.in)

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
func canVisitAllRooms(rooms [][]int) bool {
	// 2020-08-31 21:52 @bingohuang
	// 算法：1、深度优先索索
	// 复杂度：O(M+N)/O(N)
	// 效率：执行耗时:8 ms,击败了71.74% 的Go用户
	//			内存消耗:4.2 MB,击败了72.00% 的Go用户
	/*n := len(rooms)
	num := 0
	vis := make([]bool, n)
	var dfs func([][]int, int)
	dfs = func(rooms [][]int, x int) {
		vis[x] = true
		num ++
		for _, it := range rooms[x] {
			if !vis[it] {
				dfs(rooms, it)
			}
		}
	}
	dfs(rooms, 0)
	return n == num*/
	// 2020-08-31 21:56 @bingohuang
	// 算法：2、广度优先索索
	// 复杂度：O(M+N)/O(N)
	// 效率：执行耗时:4 ms,击败了98.91% 的Go用户
	//			内存消耗:4.1 MB,击败了100.00% 的Go用户
	n := len(rooms)
	num := 0
	vis := make([]bool, n)
	queue := []int{0}
	vis[0] = true
	for i := 0; i < len(queue); i++ {
		x := queue[i]
		num++
		for _, it := range rooms[x] {
			if !vis[it] {
				vis[it] = true
				queue = append(queue, it)
			}
		}
	}
	return n == num
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//有 N 个房间，开始时你位于 0 号房间。每个房间有不同的号码：0，1，2，...，N-1，并且房间里可能有一些钥匙能使你进入下一个房间。
//
// 在形式上，对于每个房间 i 都有一个钥匙列表 rooms[i]，每个钥匙 rooms[i][j] 由 [0,1，...，N-1] 中的一个整数表示，其中
//N = rooms.length。 钥匙 rooms[i][j] = v 可以打开编号为 v 的房间。
//
// 最初，除 0 号房间外的其余所有房间都被锁住。
//
// 你可以自由地在房间之间来回走动。
//
// 如果能进入每个房间返回 true，否则返回 false。
//
//
//
//
// 示例 1：
//
// 输入: [[1],[2],[3],[]]
//输出: true
//解释:
//我们从 0 号房间开始，拿到钥匙 1。
//之后我们去 1 号房间，拿到钥匙 2。
//然后我们去 2 号房间，拿到钥匙 3。
//最后我们去了 3 号房间。
//由于我们能够进入每个房间，我们返回 true。
//
//
// 示例 2：
//
// 输入：[[1,3],[3,0,1],[2],[0]]
//输出：false
//解释：我们不能进入 2 号房间。
//
//
// 提示：
//
//
// 1 <= rooms.length <= 1000
// 0 <= rooms[i].length <= 1000
// 所有房间中的钥匙数量总计不超过 3000。
//
// Related Topics 深度优先搜索 图
// 👍 142 👎 0
