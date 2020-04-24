// github.com/bingohuang/go-codes
/**
题目: 207. 课程表
难度: 2
地址: https://leetcode-cn.com/problems/course-schedule/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO207 struct {
	in1 int
	in2 [][]int
	out bool
}

func Test207(t *testing.T) {
	// add test cases
	tc := map[string]IO207{
		"1": {2, [][]int{{1, 0}}, true},
		"2": {2, [][]int{{1, 0}, {0, 1}}, false},
	}

	for k, v := range tc {
		// algo func
		out := canFinish(v.in1, v.in2)

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
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 20200424
	// BFS
	// 执行耗时:20 ms,击败了45.76% 的Go用户
	// 内存消耗:5.7 MB,击败了100.00% 的Go用户
	//indegrees := make([]int, numCourses) // 记录每门课程前置应修的课程数
	//for _, req := range prerequisites {
	//	indegrees[req[0]]++
	//}
	//queue := list.New() // 装入度为0的课程，即没有依赖可直接修的课程
	//for i := 0; i < numCourses; i++ {
	//	if indegrees[i] == 0 {
	//		queue.PushBack(i)
	//	}
	//}
	//for queue.Len() > 0 {
	//	course := queue.Remove(queue.Front()).(int)
	//	numCourses-- // 修course这门课
	//	for _, req := range prerequisites {
	//		if req[1] != course {
	//			continue
	//		}
	//		// course修过了，依赖course的课程也可以修了
	//		indegrees[req[0]]--
	//		// req[0]这门课的前置课程都修完了
	//		if indegrees[req[0]] == 0 {
	//			queue.PushBack(req[0])
	//		}
	//	}
	//}
	//return numCourses == 0

	// 20200424
	// BFS-优化
	// 执行耗时:12 ms,击败了93.22% 的Go用户
	// 内存消耗:5.8 MB,击败了100.00% 的Go用户
	//// 记录每门课程前置应修的课程数
	//indegrees := make([]int, numCourses)
	//// 记录每门课的后修课程
	//nexts := make([][]int, numCourses)
	//for _, req := range prerequisites {
	//	indegrees[req[0]] += 1
	//	nexts[req[1]] = append(nexts[req[1]], req[0])
	//}
	//queue := list.New()
	//// 装入度为0的课程，即没有依赖可直接修的课程
	//for i := 0; i < numCourses; i++ {
	//	if indegrees[i] == 0 {
	//		queue.PushBack(i)
	//	}
	//}
	//for queue.Len() > 0 {
	//	course := queue.Remove(queue.Front()).(int)
	//	numCourses-- // 修course这门课
	//	for _, next := range nexts[course] {
	//		// course修过了，依赖course的课程也可以修了
	//		indegrees[next]--
	//		if indegrees[next] == 0 { // 前置课程都修完了
	//			queue.PushBack(next)
	//		}
	//	}
	//}
	//return numCourses == 0

	// 20200424
	// DFS
	// 执行耗时:20 ms,击败了45.76% 的Go用户
	// 内存消耗:5.7 MB,击败了100.00% 的Go用户
	//flags := make([]int, numCourses)
	//var dfs func(course int) bool
	//dfs = func(course int) bool {
	//	if flags[course] == 1 {
	//		return false
	//	}
	//	if flags[course] == -1 {
	//		return true
	//	}
	//	flags[course] = 1
	//	for _, req := range prerequisites {
	//		if req[0] == course && !dfs(req[1]) {
	//			return false
	//		}
	//		/* 以下写法也是对的；两种写法都是对course邻居做判断
	//		if req[1] == course && !dfs(req[0], prerequisites, flags) {
	//			return false
	//		}*/
	//	}
	//	flags[course] = -1
	//	return true
	//}
	//for i := 0; i < numCourses; i++ {
	//	if !dfs(i) {
	//		return false
	//	}
	//}
	//return true

	// 20200424
	// DFS-优化
	// 执行耗时:12 ms,击败了93.22% 的Go用户
	// 内存消耗:6.1 MB,击败了100.00% 的Go用户
	neighbors := make([][]int, numCourses)
	for _, req := range prerequisites {
		// 写成 neighbors[req[1]] = append(neighbors[req[1]], req[0]) 也对，都是统计邻居
		neighbors[req[0]] = append(neighbors[req[0]], req[1])
	}
	flags := make([]int, numCourses)
	var dfs func(course int) bool
	dfs = func(course int) bool {
		if flags[course] == 1 {
			return false
		}
		if flags[course] == -1 {
			return true
		}
		flags[course] = 1
		for _, neighbor := range neighbors[course] {
			if !dfs(neighbor) {
				return false
			}
		}
		flags[course] = -1
		return true
	}
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
//
// 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
//
// 给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
//
//
//
// 示例 1:
//
// 输入: 2, [[1,0]]
//输出: true
//解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
//
// 示例 2:
//
// 输入: 2, [[1,0],[0,1]]
//输出: false
//解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
//
//
//
// 提示：
//
//
// 输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
// 你可以假定输入的先决条件中没有重复的边。
// 1 <= numCourses <= 10^5
//
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序
