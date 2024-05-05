package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var names []string
	for i := 0; i < n; i++ {
		var name string
		fmt.Scan(&name)
		names = append(names, name)
	}
	var scores []int
	for _, name := range names {
		hashmap := make(map[rune]int)
		for _, c := range name {
			if _, ok := hashmap[c]; !ok {
				hashmap[c] = 1
			} else {
				hashmap[c]++
			}
		}
		mapLen := len(hashmap)
		score, curScore := 0, 26
		for i := 0; i < mapLen; i++ {
			max := 0
			var key rune
			for k, count := range hashmap {
				if count > max {
					max = count
					key = k
				}
			}
			score += curScore * max
			curScore--
			delete(hashmap, key)
		}
		scores = append(scores, score)
	}
	for _, score := range scores {
		fmt.Printf("%d\n", score)
	}
}
