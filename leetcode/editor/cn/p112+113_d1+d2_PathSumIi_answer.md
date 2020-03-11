# Go 实现

> [112. 路径总和 - 简单](https://leetcode-cn.com/problems/path-sum/)
> [完整代码实现](https://github.com/bingohuang/go-codes/blob/master/leetcode/editor/cn/p112_d1_PathSum_test.go)

> [113. 路径总和 II - 中等](https://leetcode-cn.com/problems/path-sum/)
> [完整代码实现](https://github.com/bingohuang/go-codes/blob/master/leetcode/editor/cn/p113_d2_PathSumIi_test.go)

## 思考

## 算法: 前序遍历递归回溯

深度搜索所有从跟节点到叶节点的路径，检查各路径上所有节点的值之和是否为sum

[112. 路径总和](https://leetcode-cn.com/problems/path-sum/) 只需判断是否满足，无需存储路径结果，[113. 路径总和 II](https://leetcode-cn.com/problems/path-sum/)需要存储路径遍历结果，这里重点解析 [113. 路径总和 II](https://leetcode-cn.com/problems/path-sum/) 题的解法

### 关键问题
1. 使用何种数据结构存储遍历路径上的节点？
    - 栈，比如数组实现的栈
2. 在树的前序遍历位置做什么？后序遍历位置做什么？
    - 在前序遍历位置，将该节点值存储至路径栈中，并累加路径值
    - 在后续遍历位置，将该节点值从路径栈中弹出，路径值减去结点值
3. 如何判断一个节点为叶节点？当遍历到叶结点时应该做什么？
    - 叶子节点左右子树都为空
    - 如果是叶子节点，且当前路径值满足 sum 要求，则返回true退出，或者记录结果继续遍历

### 算法图解-TODO

### Go 实现 - [112. 路径总和 - 简单](https://leetcode-cn.com/problems/path-sum/) - 前序遍历递归
```go
// 只需判断是否满足，无需存储路径结果
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
```
### 复杂度分析
- 时间复杂度：$O(N)$
- 空间复杂度：$O(N)$

### Go 实现 - [113. 路径总和 II - 中等](https://leetcode-cn.com/problems/path-sum/)
```go
// 需存储路径遍历结果
func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	var path []int
	var pathValue int
	var preOrder func(node *TreeNode, sum int)
	preOrder = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		pathValue += node.Val
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && pathValue == sum {
			t := make([]int, len(path))
			copy(t, path)
			result = append(result, t)
		}
		preOrder(node.Left, sum)
		preOrder(node.Right, sum)
		pathValue -= node.Val
		path = path[:len(path)-1] // pop
	}
	preOrder(root, sum)
	return result
}
```
### 复杂度分析
- 时间复杂度：$O(N)$
- 空间复杂度：$O(N)$