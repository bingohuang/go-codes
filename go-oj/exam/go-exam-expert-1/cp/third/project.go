/*
	原题没权限，看别人博客吧：
	https://www.cnblogs.com/grandyang/p/6746528.html
*/

package main

import "fmt"

var n, m int
var queue [][]int
var head, tail int
var dist [][][]int
var vis [][][]bool
var path [][][]string

// const value
const (
	MaxDist  = 1000000 // MaxDist 定义不可达
	FourWays = 4       // FourWays 定义方向的数量
)

var way = [][]int{
	{-1, 0}, // Up
	{1, 0},  // Down
	{0, -1}, // Left
	{0, 1},  // Right
}

var turn = map[int]string{
	0: "u",
	1: "d",
	2: "l",
	3: "r",
}

func push(pos []int) {
	queue[tail] = pos
	tail++
}

func pop() []int {
	pos := queue[head]
	head++
	return pos
}

func empty() bool {
	return head == tail
}

func checkIsWall(x, y, n, m int, matrix [][]int) bool {
	if x < 0 || x >= n || y < 0 || y >= m { // 踩到边界，边界也是墙
		return true
	}
	return matrix[x][y] == 1
}

func update(next, now []int) {
	nx, ny, nf := next[0], next[1], next[2]
	x, y, f := now[0], now[1], now[2]
	subPath := turn[nf]
	if nf == f { // 当未转向时，不增加path的长度
		subPath = ""
	}
	if dist[nx][ny][nf] == MaxDist {
		dist[nx][ny][nf] = dist[x][y][f] + 1
		path[nx][ny][nf] = path[x][y][f] + subPath
	} else if dist[nx][ny][nf] == dist[x][y][f]+1 {
		if path[nx][ny][nf] > path[x][y][f]+subPath { // 当路径长度相同时，取字典序最小的路径
			path[nx][ny][nf] = path[x][y][f] + subPath
		}
	}
	if vis[nx][ny][nf] == true { // 已进队列
		return
	}
	vis[nx][ny][nf] = true
	push([]int{nx, ny, nf})
}

func bfs(matrix [][]int, car []int, des []int) {
	queue = make([][]int, n*m*FourWays)
	head, tail = 0, 0
	for i := 0; i < FourWays; i++ {
		x, y := car[0], car[1]
		vis[x][y][i] = true
		if !checkIsWall(x+way[i][0], y+way[i][1], n, m, matrix) { // 初始不能往墙上走
			push([]int{car[0], car[1], i})
			dist[car[0]][car[1]][i] = 0
			path[car[0]][car[1]][i] = turn[i]
		}
	}
	for !empty() {
		now := pop()
		x, y, f := now[0], now[1], now[2] // f 为朝向
		if x == des[0] && y == des[1] {
			continue
		}
		isWall := checkIsWall(x+way[f][0], y+way[f][1], n, m, matrix)
		if isWall { // 是墙时可以转向
			for i := 0; i < FourWays; i++ {
				nx := x + way[i][0]
				ny := y + way[i][1]
				if checkIsWall(nx, ny, n, m, matrix) { // 下个位置不能是墙
					continue
				}
				update([]int{nx, ny, i}, []int{x, y, f})
			}
		} else { // 不是墙不能转向
			update([]int{x + way[f][0], y + way[f][1], f}, []int{x, y, f})
		}
	}
}

func searchShortestWay(matrix [][]int, car []int, des []int) string {
	n = len(matrix)
	m = len(matrix[0])
	dist = make([][][]int, n)    // 四个方向下的最短路
	path = make([][][]string, n) // 最短路下的最小字典序路径
	vis = make([][][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([][]bool, m)
		dist[i] = make([][]int, m)
		path[i] = make([][]string, m)
		for j := 0; j < m; j++ {
			vis[i][j] = make([]bool, FourWays)
			dist[i][j] = make([]int, FourWays)
			path[i][j] = make([]string, FourWays)
			for k := 0; k < FourWays; k++ {
				dist[i][j][k] = MaxDist
			}
		}
	}

	bfs(matrix, car, des)
	minDist := MaxDist
	minPath := ""
	for i := 0; i < FourWays; i++ {
		if minDist > dist[des[0]][des[1]][i] {
			minDist = dist[des[0]][des[1]][i]
			minPath = path[des[0]][des[1]][i]
		} else if minDist == dist[des[0]][des[1]][i] && minPath > path[des[0]][des[1]][i] {
			minPath = path[des[0]][des[1]][i]
		}
	}
	if minPath == "" {
		return "impossible"
	}
	return minPath
}

func main() {
	matrix := [][]int{
		{0, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}
	car := []int{4, 3}
	des := []int{3, 0}
	ans := searchShortestWay(matrix, car, des)
	fmt.Println(ans)
}
