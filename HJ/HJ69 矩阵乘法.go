package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	matrix1 := make([][]int, x)
	for i := range matrix1 {
		matrix1[i] = make([]int, y)
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Scan(&matrix1[i][j])
		}
	}
	matrix2 := make([][]int, y)
	for i := range matrix2 {
		matrix2[i] = make([]int, z)
	}
	for i := 0; i < y; i++ {
		for j := 0; j < z; j++ {
			fmt.Scan(&matrix2[i][j])
		}
	}
	matrix := make([][]int, x)
	for i := range matrix {
		matrix[i] = make([]int, z)
	}
	for i := 0; i < x; i++ {
		for j := 0; j < z; j++ {
			for k := 0; k < y; k++ {
				tmp := matrix1[i][k] * matrix2[k][j]
				matrix[i][j] += tmp
			}
		}
	}
	for _, row := range matrix {
		for _, v := range row {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
