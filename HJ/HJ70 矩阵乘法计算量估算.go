package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	matrixs := make([][]int, n)
	for i := range matrixs {
		matrixs[i] = make([]int, 2)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&matrixs[i][j])
		}
	}
	var str string
	fmt.Scan(&str)
	stack := []rune{}
	totalCount := 0
	for _, c := range str {
		if c >= 'A' && c <= 'Z' {
			stack = append(stack, c)
		} else if c == ')' {
			row, col := matrixs[len(stack)-2][0], matrixs[len(stack)-1][1]
			count := matrixs[len(stack)-2][1]
			totalCount += row * count * col
			tmp1, tmp2 := matrixs[:len(stack)-2], matrixs[len(stack):]
			tmp := [][]int{}
			tmp = append(tmp, tmp1...)
			tmp = append(tmp, []int{row, col})
			tmp = append(tmp, tmp2...)
			matrixs = tmp
			stack = stack[:len(stack)-1]
		}
	}
	fmt.Println(totalCount)
}
