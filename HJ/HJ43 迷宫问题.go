package main

import "fmt"

// 深度优先搜索（递归+回溯）
func main() {
	var r, c int
	fmt.Scan(&r, &c)
	grid := [][]int{}
	for i := 0; i < r; i++ {
		var row []int
		for j := 0; j < c; j++ {
			var val int
			fmt.Scan(&val)
			row = append(row, val)
		}
		grid = append(grid, row)
	}

	var path [][]int
	var dfs func(x, y int) bool

	dfs = func(x, y int) bool {
		if x < 0 || x >= r || y < 0 || y >= c || grid[x][y] != 0 {
			return false
		}
		path = append(path, []int{x, y})
		grid[x][y] = 1
		if x == r-1 && y == c-1 {
			return true
		}
		if dfs(x-1, y) {
			return true
		}
		if dfs(x+1, y) {
			return true
		}
		if dfs(x, y-1) {
			return true
		}
		if dfs(x, y+1) {
			return true
		}
		path = path[:len(path)-1] // 该点之后不存在到达终点的路径，回溯
		return false
	}

	dfs(0, 0)

	for _, p := range path {
		fmt.Printf("(%d,%d)\n", p[0], p[1])
	}
}
