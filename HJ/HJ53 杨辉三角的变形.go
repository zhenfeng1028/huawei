package main

import "fmt"

// 动态规划
func main() {
	var n int
	fmt.Scan(&n)
	dp := make([][]int, n+1)
	for i := 1; i < n+1; i++ {
		dp[i] = make([]int, 2*i-1)
	}
	for i := 1; i < n+1; i++ {
		dp[i][0] = 1
		dp[i][2*i-2] = 1
		if i > 1 {
			dp[i][1] = i - 1
			dp[i][2*i-3] = i - 1
		}
	}
	for i := 3; i < n+1; i++ {
		for j := 2; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j-2] + dp[i-1][j]
		}
	}
	for i, v := range dp[n] {
		if v%2 == 0 {
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println(-1)
}

// 找规律
// n		1	2	3	4	5	6	7	8	9	10	11	……
// index	-1	-1	2	3	2	4	2	3	2	4	2	……
func method2() {
	var n int
	fmt.Scan(&n)
	if n < 3 {
		fmt.Println(-1)
	} else {
		switch n % 4 {
		case 1:
			fmt.Println(2)
		case 2:
			fmt.Println(4)
		case 3:
			fmt.Println(2)
		case 0:
			fmt.Println(3)
		}
	}
}
