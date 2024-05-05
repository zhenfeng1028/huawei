package main

import "fmt"

// 递归
// 放苹果分为两种情况，一种是有盘子为空，一种是每个盘子上都有苹果。
// 令(m,n)表示将m个苹果放入n个盘子中的摆放方法总数。
// 1.假设有一个盘子为空，则(m,n)问题转化为将m个苹果放在n-1个盘子上，即求得(m,n-1)即可
// 2.假设所有盘子都装有苹果，则每个盘子上至少有一个苹果，即最多剩下m-n个苹果，问题转化为将m-n个苹果放到n个盘子上，即求(m-n，n)
func main() {
	// var m, n int
	// fmt.Scan(&m, &n)
	// fmt.Println(f(m, n))
	method2()
}

func f(m, n int) int {
	if m < 0 || n < 0 {
		return 0
	} else if m == 1 || n == 1 {
		return 1
	} else {
		return f(m, n-1) + f(m-n, n)
	}
}

// 动态规划：dp[i][j]表示i个苹果放到j个盘子的放法
func method2() {
	var m, n int
	fmt.Scan(&m, &n)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		dp[i][1] = 1
	}
	for i := 1; i < n+1; i++ {
		dp[0][i] = 1
		dp[1][i] = 1
	}
	for i := 2; i < m+1; i++ {
		for j := 2; j < n+1; j++ {
			if i < j {
				dp[i][j] = dp[i][j-1] // 如果苹果比盘子少，那么和少一个盘子的放法是一样的
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-j][j] // 否则就是少一个盘子的放法加上每个盘子至少有一个苹果的放法
			}
		}
	}
	fmt.Println(dp[m][n])
}
