package main

import "fmt"

var n, m int
var arr [][]int

func main() {
	fmt.Scan(&n, &m)
	for i := 0; i < n; i++ {
		var row []int
		for j := 0; j < m; j++ {
			var tmp int
			fmt.Scan(&tmp)
			row = append(row, tmp)
		}
		arr = append(arr, row)
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 1 {
				search(i, j)
				count++
			}
		}
	}
	fmt.Println(count)
}

func search(x, y int) {
	arr[x][y] = 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			if i >= 0 && j >= 0 && i < n && j < m && arr[i][j] == 1 {
				search(i, j)
			}
		}
	}
}
