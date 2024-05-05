package main

import "fmt"

// dp[i][j]表示走到(i,j)有多少种走法
func main() {
	var m, n int
	fmt.Scan(&m, &n)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 0
	for i := 1; i < n+1; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < m+1; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	fmt.Println(dp[n][m])
}
