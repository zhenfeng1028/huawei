package main

import "fmt"

func main() {
	var s1, s2 string
	var k int
	fmt.Scan(&s1, &s2, &k)
	if len(s1)+k > len(s2) {
		fmt.Println(-1)
		return
	}
	hashmap := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		if _, ok := hashmap[s1[i]]; !ok {
			hashmap[s1[i]] = 1
		} else {
			hashmap[s1[i]]++
		}
	}
	for i := 0; i <= len(s2)-len(s1)-k; i++ {
		subStr := s2[i : len(s1)+k+i]
		hashmap2 := make(map[byte]int)
		for j := 0; j < len(subStr); j++ {
			if _, ok := hashmap2[subStr[j]]; !ok {
				hashmap2[subStr[j]] = 1
			} else {
				hashmap2[subStr[j]]++
			}
		}
		var valid bool = true
		for c, v := range hashmap {
			if vv, ok := hashmap2[c]; !ok {
				valid = false
				break
			} else {
				if vv < v {
					valid = false
					break
				}
			}
		}
		if valid {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
