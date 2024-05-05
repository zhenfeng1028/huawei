package main

import "fmt"

// 背包问题
// dp[i][j]表示i项工作t时长下的收入
func main() {
	var T, n int
	fmt.Scan(&T, &n)
	var tarr, warr []int
	for i := 0; i < n; i++ {
		var t, w int
		fmt.Scan(&t, &w)
		tarr = append(tarr, t)
		warr = append(warr, w)
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, T+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= T; j++ {
			if tarr[i-1] <= j {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-tarr[i-1]]+warr[i-1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	fmt.Println(dp[n][T])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
