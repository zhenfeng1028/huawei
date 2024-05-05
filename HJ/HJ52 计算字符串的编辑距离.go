package main

import "fmt"

// 动态规划：dp[i][j]表示word1[:i]到word2[:j]的编辑距离
func main() {
	var word1, word2 string
	fmt.Scan(&word1, &word2)
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// 边界条件初始化
	for i := 0; i < m+1; i++ {
		dp[i][0] = i // word1[:i]变成word2[:0]，删掉word1[:i]，需要i部操作
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i // word1[:0]变成word2[:i]，插入word2[:i]，需要i部操作
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 1、从word1[:i-1]编辑为word2[:j]，再删除word1[i-1]
				// 2、从word1[:i]编辑为word2[:j-1]，再插入word2[j-1]
				// 3、从word1[:i-1]编辑为word2[:j-1]，再将word1[i-1]修改为word2[j-1]
				dp[i][j] = Min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	fmt.Println(dp[m][n])
}

func Min(args ...int) int {
	min := args[0]
	for _, v := range args {
		if v < min {
			min = v
		}
	}
	return min
}
