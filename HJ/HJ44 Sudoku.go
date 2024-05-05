package main

import (
	"fmt"
)

func main() {
	grid := [][]int{}
	for i := 0; i < 9; i++ {
		var row []int
		for j := 0; j < 9; j++ {
			var val int
			fmt.Scan(&val)
			row = append(row, val)
		}
		grid = append(grid, row)
	}

	var findAnswer bool
	var dfs func(r, c int)
	var check func(r, c int, val int) bool

	check = func(r, c int, val int) bool {
		// 同一行
		for i := 0; i < 9; i++ {
			if grid[r][i] == val {
				return false
			}
		}
		// 同一列
		for i := 0; i < 9; i++ {
			if grid[i][c] == val {
				return false
			}
		}
		// 同一个九宫格
		rowLimit := r/3*3 + 3
		colLimit := c/3*3 + 3
		for i := rowLimit - 3; i < rowLimit; i++ {
			for j := colLimit - 3; j < colLimit; j++ {
				if grid[i][j] == val {
					return false
				}
			}
		}
		return true
	}

	dfs = func(r, c int) {
		if c == 9 {
			r++
			c = 0
		}
		if r == 9 {
			findAnswer = true
			return
		}
		if grid[r][c] == 0 {
			for i := 1; i <= 9; i++ {
				if check(r, c, i) { // 检查是否满足
					grid[r][c] = i
					dfs(r, c+1)
					if findAnswer { // 找到答案，直接返回
						return
					}
					grid[r][c] = 0 // 没找到答案，回溯
				}
			}
		} else {
			dfs(r, c+1)
		}
	}

	dfs(0, 0)

	for _, row := range grid {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Printf("\n")
	}
}
