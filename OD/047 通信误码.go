package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	hashmap := make(map[int][]int)
	for i := 0; i < n; i++ {
		if _, ok := hashmap[arr[i]]; !ok {
			hashmap[arr[i]] = make([]int, 0)
		}
		hashmap[arr[i]] = append(hashmap[arr[i]], i)
	}
	maxFreq := 0
	for _, v := range hashmap {
		if len(v) > maxFreq {
			maxFreq = len(v)
		}
	}
	indexGroup := [][]int{}
	for _, v := range hashmap {
		if len(v) == maxFreq {
			indexGroup = append(indexGroup, v)
		}
	}
	minLen := math.MaxInt64
	for _, v := range indexGroup {
		if v[len(v)-1]-v[0] < minLen {
			minLen = v[len(v)-1] - v[0]
		}
	}
	fmt.Println(minLen + 1)
}
