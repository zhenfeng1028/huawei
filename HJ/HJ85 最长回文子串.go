package main

import "fmt"

var str string

// 中心扩散法
func main() {
	fmt.Scan(&str)
	maxLen := 0
	for i := 0; i < len(str)-1; i++ {
		oddLen := expandFromCenter(i, i)
		evenLen := expandFromCenter(i, i+1)
		curLen := max(oddLen, evenLen)
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	fmt.Println(maxLen)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func expandFromCenter(left, right int) int {
	for left >= 0 && right < len(str) {
		if str[left] == str[right] {
			left--
			right++
		} else {
			break
		}
	}
	return right - left - 1
}
