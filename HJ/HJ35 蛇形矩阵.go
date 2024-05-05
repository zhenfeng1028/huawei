package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	matrix := [100][100]int{}
	matrix[0][0] = 1
	for i := 1; i < N; i++ {
		matrix[i][0] = matrix[i-1][0] + i
	}
	for i := 0; i < N; i++ {
		for j := 1; j < N-i; j++ {
			matrix[i][j] = matrix[i][j-1] + i + j + 1
		}
	}
	for _, row := range matrix {
		for _, num := range row {
			if num == 0 {
				break
			}
			fmt.Printf("%d ", num)
		}
		fmt.Println()
		if row[0] == 0 {
			break
		}
	}
}
