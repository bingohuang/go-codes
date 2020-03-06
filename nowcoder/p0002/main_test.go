package main

import (
	"reflect"
	"testing"
)

type IO struct {
	in  []int
	out int
}

func TestAlgo(t *testing.T) {
	// add test cases for yourself
	tc := map[string]IO{
		"1": IO{[]int{1, 2, 2}, 4},
		"2": IO{[]int{1, 2, 3}, 6},
		"3": IO{[]int{2, 2, 2}, 3},
		"4": IO{[]int{2, 2, 1}, 4},
		"5": IO{[]int{3, 2, 1}, 6},
		"6": IO{[]int{3, 1, 2}, 5},
		"7": IO{[]int{1, 4, 3, 2}, 7},
	}

	for k, v := range tc {
		out := Algo(v.in)
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}
}

// 标题：分子弹 | 时间限制：1秒 | 内存限制：32768K | 语言限制： 不限
//射击训练，需要给每个士兵发子弹，发子弹的个数规则是根据士兵的训练成绩给定的，传入士兵的训练成绩，
// 要求相邻士兵中，成绩好的士兵的子弹数必须更多，每个士兵至少分配一个子弹。
// 输入描述：
// 输入每个士兵的训练成绩，如1,2,3，代表3位士兵的成绩分别为1,2,3
// 输出描述：
// 最少分配的子弹数
// 示例：输入：1,2,2，输出：4
func Algo(in []int) (out int) {
	// 先排个序,不能排
	//sort.Ints(in)
	// 1 4 3 2
	// 1
	// 1 2
	// 1 2 1
	// 1 3 2 1

	// 如果士兵成绩是从低往高有顺序
	r := 0
	for i := 0; i < len(in); i++ {
		if i == 0 {
			r = 1
			out += r
			continue
		}
		if in[i] > in[i-1] {
			r = r + 1
			out += r
		} else if in[i] == in[i-1] {
			r = 1
			out += r
		}
	}

	// 如果成绩是没有顺序的（大概率），就要考虑回溯了

	return out
}
