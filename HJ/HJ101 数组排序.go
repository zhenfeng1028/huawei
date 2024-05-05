package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, order int
	var arr []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		arr = append(arr, tmp)
	}
	fmt.Scan(&order)
	sort.Ints(arr)
	if order == 0 {
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", arr[i])
		}
	} else {
		for i := n - 1; i >= 0; i-- {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println()
}
