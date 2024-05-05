package main

import "fmt"

var n int     // 充电设备的个数
var arr []int // 每个充电设备的输出功率
var pMax int  // 充电站最大输出功率
var res int   // 最优元素

func main() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var p int
		fmt.Scan(&p)
		arr = append(arr, p)
	}
	fmt.Scan(&pMax)
	dfs(0, 0)
	fmt.Println(res)
}

func dfs(level int, cur int) {
	if cur > pMax {
		return
	}
	res = max(res, cur)
	for i := level; i < n; i++ {
		cur += arr[i]
		dfs(level+1, cur)
		cur -= arr[i]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 动态规划：01背包问题
// dp[i][j]表示前i个充电设备在最大输出功率为j时的最优元素
func method2() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var p int
		fmt.Scan(&p)
		arr = append(arr, p)
	}
	fmt.Scan(&pMax)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, pMax+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= pMax; j++ {
			dp[i][j] = dp[i-1][j] // 默认不将当前设备的输出功率放进去
			if j >= arr[i-1] {    // 当前最大功率能否承受当前设备的输出功率
				dp[i][j] = max(dp[i][j], dp[i-1][j-arr[i-1]]+arr[i-1])
			}
		}
	}
	fmt.Println(dp)
	if dp[n][pMax] > pMax {
		fmt.Println(0)
	}
	fmt.Println(dp[n][pMax])
}
