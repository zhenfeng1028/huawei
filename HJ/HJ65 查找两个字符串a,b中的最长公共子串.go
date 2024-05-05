package main

import (
	"fmt"
	"strings"
)

// 暴力枚举
func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	tmp := str1
	if len(str1) > len(str2) {
		str1 = str2
		str2 = tmp
	}
	maxLen := 0
	maxSubStr := ""
	for i := 0; i < len(str1); i++ {
		for j := i + 1; j <= len(str1); j++ {
			if strings.Contains(str2, str1[i:j]) && len(str1[i:j]) > maxLen {
				maxLen = len(str1[i:j])
				maxSubStr = str1[i:j]
			}
		}
	}
	fmt.Println(maxSubStr)
}

// 枚举改进
// 其实找子串不用像方法一一样完全枚举，我们完全可以遍历两个字符串的所有字符串作为起始，
// 然后同时开始检查字符是否相等，相等则不断后移，增加子串长度，如果不等说明以这两个为起点的子串截止了，
// 不会再有了，后续比较长度维护最大值即可。
func method2() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	tmp := str1
	if len(str1) > len(str2) {
		str1 = str2
		str2 = tmp
	}
	maxLen := 0
	maxSubStr := ""
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			x, y := i, j
			length := 0
			for x < len(str1) && y < len(str2) && str1[x] == str2[y] {
				x++
				y++
				length++
			}
			if length > maxLen {
				maxLen = length
				maxSubStr = str1[i:x]
			}
		}
	}
	fmt.Println(maxSubStr)
}

// 动态规划
// 我们可以用dp[i][j]表示在str1中以第i个字符结尾在str2中以第j个字符结尾时的公共子串长度，
// 遍历两个字符串填充dp数组，在这个过程中比较维护最大值即可。

// 转移方程为：如果遍历到的该位两个字符相等，则此时长度等于两个前一位长度+1，dp[i][j]=dp[i−1][j−1]+1，
// 如果遍历到该位时两个字符不相等，则置为0，因为这是子串，必须连续相等，断开要重新开始。
func method3() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	tmp := str1
	if len(str1) > len(str2) {
		str1 = str2
		str2 = tmp
	}
	dp := make([][]int, len(str1)+1)
	for i := range dp {
		dp[i] = make([]int, len(str2)+1)
	}
	maxLen := 0
	begin := 0 // 最长公共子串的起始下标
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1 // 前一个长度加一
			} else {
				dp[i][j] = 0 // 如果第i个字符和第j个字符不同，则以它们结尾的子串不可能相同
			}
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
				begin = i - maxLen
			}
		}
	}
	fmt.Println(str1[begin : begin+maxLen])
}
