// github.com/bingohuang/go-codes
/**
题目: 729. 我的日程安排表 I
难度: 2
地址: https://leetcode-cn.com/problems/my-calendar-i/
*/
package test

import (
	"container/list"
	"math"
	"testing"
)

// input and ouput
type IO729 struct {
	in  []string
	out []int
}

func Test729(t *testing.T) {
	// add test cases

	myCal := Constructor729()
	res := myCal.Book(10, 20)
	if res != true {
		t.Errorf("case-1: except answer: [true], get answer: [%v]", res)
	}
	res = myCal.Book(15, 25)
	if res != false {
		t.Errorf("case-2: except answer: [false], get answer: [%v]", res)
	}
	res = myCal.Book(20, 30)
	if res != true {
		t.Errorf("case-1: except answer: [true], get answer: [%v]", res)
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
// 20200422
// 执行耗时:212 ms,击败了18.18% 的Go用户
// 内存消耗:6.7 MB,击败了100.00% 的Go用户
type Interval struct {
	start, end int
}
type MyCalendar struct {
	calendar *list.List
}

func Constructor729() MyCalendar {
	return MyCalendar{calendar: list.New()}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for e := this.calendar.Front(); e != nil; e = e.Next() {
		interval := e.Value.(Interval)
		if max(start, interval.start) < min(end, interval.end) {
			return false
		}
	}
	this.calendar.PushBack(Interval{
		start: start,
		end:   end,
	})
	return true
}

//func max(a, b int) int {
//	return int(math.Max(float64(a), float64(b)))
//}
//
//func min(a, b int) int {
//	return int(math.Min(float64(a), float64(b)))
//}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内没有其他安排，则可以存储这个新的日程安排。
//
// MyCalendar 有一个 book(int start, int end)方法。它意味着在 start 到 end 时间内增加一个日程安排，注意，这里
//的时间是半开区间，即 [start, end), 实数 x 的范围为， start <= x < end。
//
// 当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生重复预订。
//
// 每次调用 MyCalendar.book方法时，如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true。否则，返回 false 并且不要将该
//日程安排添加到日历中。
//
// 请按照以下步骤调用 MyCalendar 类: MyCalendar cal = new MyCalendar(); MyCalendar.book(st
//art, end)
//
// 示例 1:
//
// MyCalendar();
//MyCalendar.book(10, 20); // returns true
//MyCalendar.book(15, 25); // returns false
//MyCalendar.book(20, 30); // returns true
//解释:
//第一个日程安排可以添加到日历中.  第二个日程安排不能添加到日历中，因为时间 15 已经被第一个日程安排预定了。
//第三个日程安排可以添加到日历中，因为第一个日程安排并不包含时间 20 。
//
//
// 说明:
//
//
// 每个测试用例，调用 MyCalendar.book 函数最多不超过 100次。
// 调用函数 MyCalendar.book(start, end)时， start 和 end 的取值范围为 [0, 10^9]。
//
// Related Topics 数组
