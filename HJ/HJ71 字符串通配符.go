package main

import (
	"fmt"
	"strings"
)

// 递归
func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	fmt.Println(match(strings.ToLower(str1), strings.ToLower(str2)))
}

func match(pattern, origin string) bool {
	if len(pattern) == 0 && len(origin) == 0 {
		return true
	}
	if len(pattern) == 0 || len(origin) == 0 {
		return false
	}
	if pattern[0] == origin[0] {
		return match(pattern[1:], origin[1:])
	}
	if pattern[0] == '?' {
		if origin[0] >= '0' && origin[0] <= '9' || origin[0] >= 'a' && origin[0] <= 'z' {
			return match(pattern[1:], origin[1:])
		}
	}
	if pattern[0] == '*' {
		i := 0
		for i < len(pattern) && pattern[i] == '*' { // 考虑连续*的情况，简化为一个*
			i++
		}
		// 遇到*号，匹配0个，匹配1个，匹配多个（pattern不动，origin往后挪一位）
		return match(pattern[i:], origin) || match(pattern[i:], origin[1:]) || match(pattern[i-1:], origin[1:])
	}
	return false
}
