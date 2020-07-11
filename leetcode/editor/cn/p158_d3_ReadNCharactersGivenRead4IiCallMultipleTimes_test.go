// github.com/bingohuang/go-codes
/**
题目: 158. 用 Read4 读取 N 个字符 II
难度: 3
地址: https://leetcode-cn.com/problems/read-n-characters-given-read4-ii-call-multiple-times/
*/
package test

import (
	"fmt"
	"reflect"
	"testing"
)

// input and ouput
type IO158 struct {
	in  []string
	out []int
}

func Test158(t *testing.T) {
	// add test cases
	tc := map[string]IO158{
		"1": {[]string{}, []int{}},
	}

	for k, v := range tc {
		// algo func
		out := p158(v.in)

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
/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {

	}
}

//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//给你一个文件，并且该文件只能通过给定的 read4 方法来读取，请实现一个方法使其能够读取 n 个字符。注意：你的 read 方法可能会被调用多次。
//
// read4 的定义：
//
// 参数类型: char[] buf
//返回类型: int
//
//注意: buf[] 是目标缓存区不是源缓存区，read4 的返回结果将会复制到 buf[] 当中。
//
//
// 下列是一些使用 read4 的例子：
//
// File file("abcdefghijk"); // 文件名为 "abcdefghijk"， 初始文件指针 (fp) 指向 'a'
//char[] buf = new char[4]; // 创建一个缓存区使其能容纳足够的字符
//read4(buf); // read4 返回 4。现在 buf = "abcd"，fp 指向 'e'
//read4(buf); // read4 返回 4。现在 buf = "efgh"，fp 指向 'i'
//read4(buf); // read4 返回 3。现在 buf = "ijk"，fp 指向文件末尾
//
// read 方法：
//
// 通过使用 read4 方法，实现 read 方法。该方法可以从文件中读取 n 个字符并将其存储到缓存数组 buf 中。您 不能 直接操作文件。
//
// 返回值为实际读取的字符。
//
// read 的定义：
//
// 参数:   char[] buf, int n
//返回值: int
//
//注意: buf[] 是目标缓存区不是源缓存区，你需要将结果写入 buf[] 中。
//
//
//
//
// 示例 1：
//
// File file("abc");
//Solution sol;
//// 假定 buf 已经被分配了内存，并且有足够的空间来存储文件中的所有字符。
//sol.read(buf, 1); // 当调用了您的 read 方法后，buf 需要包含 "a"。 一共读取 1 个字符，因此返回 1。
//sol.read(buf, 2); // 现在 buf 需要包含 "bc"。一共读取 2 个字符，因此返回 2。
//sol.read(buf, 1); // 由于已经到达了文件末尾，没有更多的字符可以读取，因此返回 0。
//
//
// Example 2:
//
// File file("abc");
//Solution sol;
//sol.read(buf, 4); // 当调用了您的 read 方法后，buf 需要包含 "abc"。 一共只能读取 3 个字符，因此返回 3。
//sol.read(buf, 1); // 由于已经到达了文件末尾，没有更多的字符可以读取，因此返回 0。
//
//
// 注意：
//
//
// 你 不能 直接操作该文件，文件只能通过 read4 获取而 不能 通过 read。
// read 函数可以被调用 多次。
// 请记得 重置 在 Solution 中声明的类变量（静态变量），因为类变量会 在多个测试用例中保持不变，影响判题准确。请 查阅 这里。
// 你可以假定目标缓存数组 buf 保证有足够的空间存下 n 个字符。
// 保证在一个给定测试用例中，read 函数使用的是同一个 buf。
//
// Related Topics 字符串
