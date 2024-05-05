package main

import "fmt"

var m, n int
var grid = [][]string{}
var entry [2]int   // 入口坐标
var entryCount int // 入口个数

func main() {
	fmt.Scan(&m, &n)
	for i := 0; i < m; i++ {
		var row []string
		for j := 0; j < n; j++ {
			var val string
			fmt.Scan(&val)
			row = append(row, val)
		}
		grid = append(grid, row)
	}
	maxArea := 0
	res := make([]int, 3) // 存放入口坐标及区域大小
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == "O" {
				region := dfs(i, j)
				if entryCount == 1 && maxArea < len(region) {
					res[0] = entry[0]
					res[1] = entry[1]
					res[2] = len(region)
					maxArea = len(region)
				}
				entryCount = 0 // 重置入口数量
			}
		}
	}
	if res[2] > 0 {
		fmt.Printf("%d %d %d\n", res[0], res[1], res[2])
	} else if maxArea != 0 {
		fmt.Println(maxArea)
	} else {
		fmt.Println("NULL")
	}
}

func dfs(x, y int) [][2]int {
	if x < 0 || y < 0 || x >= m || y >= n || grid[x][y] != "O" {
		return nil
	}
	res := [][2]int{}
	res = append(res, [2]int{x, y})
	grid[x][y] = "X"
	if x == 0 || x == m-1 || y == 0 || y == n-1 {
		entry[0] = x
		entry[1] = y
		entryCount++
	}
	tmp1 := dfs(x-1, y)
	if tmp1 != nil {
		for _, v := range tmp1 {
			res = append(res, v)
		}
	}
	tmp2 := dfs(x+1, y)
	if tmp2 != nil {
		for _, v := range tmp2 {
			res = append(res, v)
		}
	}
	tmp3 := dfs(x, y-1)
	if tmp3 != nil {
		for _, v := range tmp3 {
			res = append(res, v)
		}
	}
	tmp4 := dfs(x, y+1)
	if tmp4 != nil {
		for _, v := range tmp4 {
			res = append(res, v)
		}
	}
	return res
}
