package main

import "fmt"

// 暴力解法
func main() {
	var input string
	fmt.Scan(&input)
	perfectCount := len(input) / 4
	minLen := len(input)
	// 原字符串各个字符的个数
	hashmap := map[rune]int{
		'A': 0,
		'B': 0,
		'C': 0,
		'D': 0,
	}
	for _, c := range input {
		hashmap[c]++
	}
	if hashmap['A'] == perfectCount && hashmap['B'] == perfectCount && hashmap['C'] == perfectCount {
		fmt.Println(0)
		return
	}
	for i := 0; i < len(input); i++ {
		for j := i + 1; j <= len(input); j++ {
			substr := input[i:j]
			hashmap2 := map[rune]int{
				'A': 0,
				'B': 0,
				'C': 0,
				'D': 0,
			}
			for _, c := range substr {
				hashmap2[c]++
			}
			// 当子串两侧各个字符小于等于perfectCount时，该子串可更换
			if hashmap['A']-hashmap2['A'] <= perfectCount && hashmap['B']-hashmap2['B'] <= perfectCount &&
				hashmap['C']-hashmap2['C'] <= perfectCount && hashmap['D']-hashmap2['D'] <= perfectCount {
				if j-i < minLen {
					minLen = j - i
				}
			}
		}
	}
	fmt.Println(minLen)
}
