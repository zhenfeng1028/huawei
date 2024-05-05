package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	if 1 <= m && m <= 9 && 1 <= n && n <= 9 {
		fmt.Println(0)
	} else {
		fmt.Println(-1)
	}
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)
	if 0 <= x1 && x1 < m && 0 <= y1 && y1 < n && 0 <= x2 && x2 < m && 0 <= y2 && y2 < n {
		fmt.Println(0)
	} else {
		fmt.Println(-1)
	}
	var x int
	fmt.Scan(&x)
	if x < m && m <= 8 {
		fmt.Println(0)
	} else {
		fmt.Println(-1)
	}
	var y int
	fmt.Scan(&y)
	if y < n && n <= 8 {
		fmt.Println(0)
	} else {
		fmt.Println(-1)
	}
	fmt.Scan(&x, &y)
	if x <= m-1 && y <= n-1 {
		fmt.Println(0)
	} else {
		fmt.Println(-1)
	}
}
