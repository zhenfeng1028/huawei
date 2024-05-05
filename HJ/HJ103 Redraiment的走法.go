package main

import (
	"fmt"
	"sort"
)

// 动态规划：dp[i]表示终点为第i个桩的最大次数
func main() {
	var n int
	var arr []int
	fmt.Scan(&n)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		arr = append(arr, tmp)
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	sort.Ints(dp)
	fmt.Println(dp[len(dp)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
