package main

import (
	"fmt"
	"strings"
)

// 暴力解法
func main() {
	var text1, text2 string
	fmt.Scan(&text1, &text2)
	maxLen := 0
	maxSubStr := ""
	for i := 0; i < len(text1); i++ {
		for j := i + 1; j <= len(text1); j++ {
			if strings.Contains(text2, text1[i:j]) && j-i > maxLen {
				maxLen = j - i
				maxSubStr = text1[i:j]
			}
		}
	}
	fmt.Println(maxSubStr)
}

// 动态规划
func method2() {
	var text1, text2 string
	fmt.Scan(&text1, &text2)
	if len(text1) > len(text2) {
		tmp := text1
		text1 = text2
		text2 = tmp
	}
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	maxLen := 0
	begin := 0
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] != text2[j-1] {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
				begin = i - maxLen
			}
		}
	}
	fmt.Println(text1[begin : begin+maxLen])
}
