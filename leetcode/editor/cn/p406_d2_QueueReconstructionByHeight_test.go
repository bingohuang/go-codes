// github.com/bingohuang/go-codes
/**
题目: 406. 根据身高重建队列
难度: 2
地址: https://leetcode-cn.com/problems/queue-reconstruction-by-height/
*/
package test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// input and ouput
type IO406 struct {
	in  [][]int
	out [][]int
}

func Test406(t *testing.T) {
	// add test cases
	tc := map[string]IO406{
		"1": {[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}, [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}},
		"2": {[][]int{{8, 0}, {4, 4}, {8, 1}, {5, 0}, {6, 1}, {5, 2}}, [][]int{{5, 0}, {8, 0}, {5, 2}, {6, 1}, {4, 4}, {8, 1}}},
	}

	for k, v := range tc {
		// algo func
		out := reconstructQueue(v.in)

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
func reconstructQueue(people [][]int) [][]int {
	// 20200423
	// 执行耗时:12 ms,击败了96.85% 的Go用户
	// 内存消耗:5.9 MB,击败了66.67% 的Go用户
	// 执行用时 : 16 ms
	// 内存消耗 : 5.9 MB
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}

		return people[i][0] > people[j][0]
	})
	res := make([][]int, 0)
	for _, p := range people {
		k := p[1]
		if k >= len(res) {
			res = append(res, p)
		} else {
			res = append(res, []int{})
			copy(res[k+1:], res[k:len(res)-1])
			res[k] = p
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//假设有打乱顺序的一群人站成一个队列。 每个人由一个整数对(h, k)表示，其中h是这个人的身高，k是排在这个人前面且身高大于或等于h的人数。 编写一个算法来
//重建这个队列。
//
// 注意：
//总人数少于1100人。
//
// 示例
//
//
//输入:
//[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
//
//输出:
//[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
//
// Related Topics 贪心算法
