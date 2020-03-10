# Go 实现

> [51. N皇后 - 困难](https://leetcode-cn.com/problems/n-queens/)

> [完整代码实现](https://github.com/bingohuang/go-codes/blob/master/leetcode/editor/cn/p51_d3_NQueens_test.go)

## 思考-TODO
- 皇后的攻击范围

- 棋盘与皇后表示


## 算法 1: 回溯+剪枝算法（比较耗时耗空间）
[参考题解](https://leetcode-cn.com/problems/n-queens/solution/golanghui-su-by-sealyun/)

```go
for j in 列：
    如果不合法
        continue
    在此行j的地方放皇后
    递归
    回溯
```

### Go 实现1
```go
func solveNQueens1(n int) [][]string {
	res = [][]string{}
	chessBoard := make([][]bool, n)
	for i := 0; i < n; i++ {
		chessBoard[i] = make([]bool, n)
	}
	trackBack1(chessBoard, [][]byte{})
	return res
}

func trackBack1(chessBoard [][]bool, track [][]byte) {
	if len(track) == len(chessBoard) {
		t := make([]string, len(track))
		for k, bs := range track {
			t[k] = string(bs)
		}
		res = append(res, t)
	}

	for j := 0; j < len(chessBoard); j++ {
		if !valid(chessBoard, len(track), j) {
			continue
		}
		bs := getLine(len(chessBoard))
		bs[j] = 'Q'
		chessBoard[len(track)][j] = true
		track = append(track, bs)
		trackBack1(chessBoard, track)
		track = track[:len(track)-1]
		chessBoard[len(track)][j] = false
	}
}

func valid(chessBoard [][]bool, row, cow int) bool {
	var i, j int
	for i = 0; i < row; i++ {
		if chessBoard[i][cow] == true {
			return false
		}
	}

	for j = 0; j < len(chessBoard); j++ {
		if chessBoard[row][j] == true {
			return false
		}
	}

	i, j = row, cow
	for i >= 0 && j >= 0 {
		if chessBoard[i][j] == true {
			return false
		}
		i--
		j--
	}

	i, j = row, cow
	for i >= 0 && j < len(chessBoard) {
		if chessBoard[i][j] == true {
			return false
		}
		j++
		i--
	}
	return true
}

func getLine(n int) []byte {
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = '.'
	}
	return bs
}
```
### 复杂度分析
- 时间复杂度：$O(N^4)$
- 空间复杂度：$O(N^2)$

## 算法2：优雅的解法-TODO


### 复杂度分析
- 时间复杂度：$O(N^3)$
- 空间复杂度：$O(N)$

## 算法3：最优的解法-TODO

### 复杂度分析
- 时间复杂度：$O(N^2)$
- 空间复杂度：$O(N)$