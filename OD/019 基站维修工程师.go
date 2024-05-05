package main

import (
	"fmt"
	"math"
)

var n int             // 基站个数
var distances [][]int // 基站之间的距离
var visited []bool    // 基站是否被访问
var minDistance int   // 最短距离

func main() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var row []int
		for j := 0; j < n; j++ {
			var tmp int
			fmt.Scan(&tmp)
			row = append(row, tmp)
		}
		distances = append(distances, row)
	}
	visited = make([]bool, n)
	minDistance = math.MaxInt
	dfs(0, 0, 0)
	fmt.Println(minDistance)
}

func dfs(curPos int, step int, sum int) {
	if step == n-1 {
		sum += distances[curPos][0] // 从当前位置回到基站1
		minDistance = min(minDistance, sum)
		return
	}
	for i := 1; i < n; i++ {
		if !visited[i] {
			visited[i] = true
			dfs(i, step+1, sum+distances[curPos][i])
			visited[i] = false
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
