package main

import (
	"fmt"
	"sort"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	var weights []int
	for i := 0; i < n; i++ {
		var weight int
		fmt.Scan(&weight)
		weights = append(weights, weight)
	}
	sort.Ints(weights)
	left, right := 0, len(weights)-1
	count := 0
	for left < right {
		if weights[left]+weights[right] > m {
			count++
			right--
			continue
		}
		count++
		left++
		right--
	}
	if left == right {
		count++
	}
	fmt.Println(count)
}
