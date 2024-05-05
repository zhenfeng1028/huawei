package main

import (
	"fmt"
	"math"
)

// 同最长回文子串
func main() {
	var password string
	fmt.Scan(&password)
	fmt.Println(longestPalindrome(password))
}

// 中心扩散法
func longestPalindrome(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	maxLen := math.MinInt
	for i := 0; i < len(s)-1; i++ {
		oddLen := expandFromCenter(s, i, i)
		evenLen := expandFromCenter(s, i, i+1)
		curLen := max(oddLen, evenLen)
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func expandFromCenter(s string, left, right int) int {
	for left >= 0 && right <= len(s)-1 {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return right - left - 1
}

// 动态规划
func longestPalindrome2(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	maxLen := -1
	dp := [2500][2500]bool{}
	// 单个字符一定是对称的
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}
	// 从上往下，从左往右
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] == true && j-i+1 > maxLen {
				maxLen = j - i + 1
			}
		}
	}
	return maxLen
}
