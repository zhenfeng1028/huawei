package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	locs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&locs[i])
	}
	sort.Ints(locs)
	fmt.Println(locs)
	if len(locs)%2 == 0 {
		fmt.Println(locs[len(locs)/2-1])
	} else {
		fmt.Println(locs[len(locs)/2])
	}
}
