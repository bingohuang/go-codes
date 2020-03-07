# Go 实现(含图解)

> [206. 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)

> [完整代码实现](https://github.com/bingohuang/go-codes/blob/master/leetcode/editor/cn/p206_ReverseLinkedList_test.go)

## 思路
- 要求: 已知链表头节点指针head，将链表逆序。(不可申请额外空间)

![](https://pic.leetcode-cn.com/cbdb66b5483a7781ce1cfd4b458f8387fbcfd909704109b3b71e58b4a5bab689.png)

- 整体思路

![](https://pic.leetcode-cn.com/724e40955206daea0b5887bc77d3325ee5a05f9f3a59b2c798880a7bb62e2865.png)

- 关键步骤

![](https://pic.leetcode-cn.com/43cf406296558661126dd04fadac55545d93007ebfa6ce6b37f574b4fdff10f4.png)

![](https://pic.leetcode-cn.com/17337d26cf2c714249b5f2be4419a5095a201049b5b5d156c24d867a75e3f36a.png)

![](https://pic.leetcode-cn.com/415e0d49cb49faf03f53f3c5f3e88e9fd7cf187177dc9c7a02451350e1c8cd92.png)

- 效果

![](https://pic.leetcode-cn.com/78e9b6ecd1a8f371b971e7a42673b297606a39bd4de56fa3f2401c2b5d198699.png)

## 算法1. 迭代法
```go
// 算法1: 翻转链表
func reverseList1(head *ListNode) *ListNode {
	var newHead, next *ListNode

	for head != nil {
		next = head.Next
		head.Next = newHead
		newHead = head
		head = next
	}

	return newHead
}
```

### 复杂度分析
- 时间复杂度：O(N)
- 空间复杂度：O(1)