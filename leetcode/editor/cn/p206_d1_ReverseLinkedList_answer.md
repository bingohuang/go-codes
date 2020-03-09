# Go 实现(含图解)

> [206. 反转链表 - 简单](https://leetcode-cn.com/problems/reverse-linked-list/)

> [完整代码实现](https://github.com/bingohuang/go-codes/blob/master/leetcode/editor/cn/p206_d1_ReverseLinkedList_test.go)

## 思路
- 要求: 已知链表头节点指针head，将链表逆序。(不可申请额外空间)

![](https://leetcode-cn.oss-cn-hangzhou.aliyuncs.com/p206/p206-1.png)

- 整体思路

![](https://leetcode-cn.oss-cn-hangzhou.aliyuncs.com/p206/p206-2.png)

- 关键步骤

![](https://leetcode-cn.oss-cn-hangzhou.aliyuncs.com/p206/p206-9.png)

![](https://leetcode-cn.oss-cn-hangzhou.aliyuncs.com/p206/p206-10.png)

![](https://leetcode-cn.oss-cn-hangzhou.aliyuncs.com/p206/p206-11.png)

- 效果

![](https://leetcode-cn.oss-cn-hangzhou.aliyuncs.com/p206/p206-12.png)

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
- 时间复杂度：$O(N)$
- 空间复杂度：$O(1)$