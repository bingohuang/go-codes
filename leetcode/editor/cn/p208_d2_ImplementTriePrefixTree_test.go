// github.com/bingohuang/go-codes
/**
题目: 208. 实现 Trie (前缀树)
难度: 2
地址: https://leetcode-cn.com/problems/implement-trie-prefix-tree/
*/
package test

import (
	"fmt"
	"testing"
)

// input and ouput
type IO208 struct {
	in  []string
	out []int
}

func Test208(t *testing.T) {
	// add test cases
	/*tc := map[string]IO208{
		"0": {[]string{}, []int{}},
	}

	for k, v := range tc {
		// algo func
		out := p208(v.in)

		fmt.Printf("case-%v:\n", k)
		fmt.Printf("\tinput: %v\n", v.in)
		fmt.Printf("\toutput: %v\n", out)
		fmt.Printf("\texcept: %v\n", v.out)

		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("case-%v: except answer: [%v], get answer: [%v]", k, v.out, out)
		}
	}*/

	obj := Constructor208()
	obj.Insert("apple")
	fmt.Println(obj.Search("apple"))   // true
	fmt.Println(obj.Search("app"))     // false
	fmt.Println(obj.StartsWith("app")) // true

	//trie.insert("apple");
	//trie.search("apple");   // 返回 true
	//trie.search("app");     // 返回 false
	//trie.startsWith("app"); // 返回 true
	//trie.insert("app");
	//trie.search("app");     // 返回 true
}

//leetcode submit region begin(Prohibit modification and deletion)
// 20200817
// https://leetcode-cn.com/problems/implement-trie-prefix-tree/solution/golangqian-zhui-shu-ti-jie-si-lu-by-user4484c/
// 执行耗时:84 ms,击败了37.74% 的Go用户
// 内存消耗:18.1 MB,击败了49.23% 的Go用户
type Trie struct {
	name    string
	subTrie map[string]*Trie
	isWord  bool
}

/** Initialize your data structure here. */
func Constructor208() Trie {
	return Trie{
		name:    "",
		subTrie: make(map[string]*Trie, 0),
		isWord:  false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this
	for _, v := range word {
		value, ok := root.subTrie[string(v)]
		if ok {
			root = value
		} else {
			newNode := make(map[string]*Trie, 0)
			root.subTrie[string(v)] = &Trie{
				name:    string(v),
				subTrie: newNode,
				isWord:  false,
			}
			root = root.subTrie[string(v)]
		}
	}
	root.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	isWord, ok := this.searchbyWord(word)
	if ok && isWord {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	_, ok := this.searchbyWord(prefix)
	return ok
}

func (this *Trie) searchbyWord(word string) (bool, bool) {
	var (
		root  = this
		value *Trie
		ok    bool
	)
	for _, v := range word {
		value, ok = root.subTrie[string(v)]
		if ok {
			root = value
			continue
		}
		return false, false
	}
	return value.isWord, true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)

/* 题目详情 */
//实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
//
// 示例:
//
// Trie trie = new Trie();
//
//trie.insert("apple");
//trie.search("apple");   // 返回 true
//trie.search("app");     // 返回 false
//trie.startsWith("app"); // 返回 true
//trie.insert("app");
//trie.search("app");     // 返回 true
//
// 说明:
//
//
// 你可以假设所有的输入都是由小写字母 a-z 构成的。
// 保证所有输入均为非空字符串。
//
// Related Topics 设计 字典树
// 👍 374 👎 0
