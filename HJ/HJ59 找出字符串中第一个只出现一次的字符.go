package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	hashmap := make(map[rune]int)
	for _, c := range str {
		if _, ok := hashmap[c]; !ok {
			hashmap[c] = 1
		} else {
			hashmap[c]++
		}
	}
	for k, v := range hashmap {
		if v > 1 {
			delete(hashmap, k)
		}
	}
	var find bool
	for _, c := range str {
		if _, ok := hashmap[c]; ok {
			fmt.Println(string(c))
			find = true
			break
		}
	}
	if !find {
		fmt.Println(-1)
	}
}
