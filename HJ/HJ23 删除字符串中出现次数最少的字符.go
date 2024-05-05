package main

import (
	"fmt"
	"math"
)

func main() {
	var str string
	fmt.Scan(&str)
	hashmap := make(map[byte]int)
	for _, c := range str {
		if _, ok := hashmap[byte(c)]; !ok {
			hashmap[byte(c)] = 1
		} else {
			hashmap[byte(c)]++
		}
	}
	min := math.MaxInt
	for _, v := range hashmap {
		if v < min {
			min = v
		}
	}
	hashmap2 := make(map[byte]struct{})
	for k, v := range hashmap {
		if v == min {
			hashmap2[k] = struct{}{}
		}
	}
	res := ""
	for _, c := range str {
		if _, ok := hashmap2[byte(c)]; !ok {
			res += string(c)
		}
	}
	fmt.Println(res)
}
