package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	digits := []int{}
	hashmap := make(map[int]struct{})
	for num > 0 {
		digit := num % 10
		if _, ok := hashmap[digit]; !ok {
			hashmap[digit] = struct{}{}
			digits = append(digits, digit)
		}
		num /= 10
	}
	res := 0
	for _, v := range digits {
		res = res*10 + v
	}
	fmt.Println(res)
}
