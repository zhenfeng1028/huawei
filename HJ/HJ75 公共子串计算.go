package main

import (
	"fmt"
	"strings"
)

// 同第65题
func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	tmp := str1
	if len(str1) > len(str2) {
		str1 = str2
		str2 = tmp
	}
	maxLen := 0
	for i := 0; i < len(str1); i++ {
		for j := i + 1; j <= len(str1); j++ {
			if strings.Contains(str2, str1[i:j]) && len(str1[i:j]) > maxLen {
				maxLen = len(str1[i:j])
			}
		}
	}
	fmt.Println(maxLen)
}
