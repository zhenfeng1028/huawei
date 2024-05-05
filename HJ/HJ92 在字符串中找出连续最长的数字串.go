package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	maxLen := 0
	subStr := []string{}
	for i := 0; i < len(str); i++ {
		if '0' <= str[i] && str[i] <= '9' {
			k := i
			for k < len(str) {
				if '0' <= str[k] && str[k] <= '9' {
					k++
				} else {
					break
				}
			}
			if k-i > maxLen {
				maxLen = k - i
				subStr = []string{str[i:k]}
			} else if k-i == maxLen {
				subStr = append(subStr, str[i:k])
			}
		}
	}
	fmt.Printf("%s,%d\n", strings.Join(subStr, ""), maxLen)
}
