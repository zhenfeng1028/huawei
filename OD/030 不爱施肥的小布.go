package main

import (
	"fmt"
	"sort"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	var fields []int
	for i := 0; i < m; i++ {
		var field int
		fmt.Scan(&field)
		fields = append(fields, field)
	}
	sort.Ints(fields)
	if m > n {
		fmt.Println(-1)
		return
	}
	maxField := fields[len(fields)-1]
	if m == n {
		fmt.Println(maxField)
		return
	}
	left, right := 1, maxField
	for left < right {
		mid := left + (right-left)/2
		sumDays := 0
		for _, field := range fields {
			if field%mid == 0 {
				sumDays += field / mid
			} else {
				sumDays += field/mid + 1
			}
		}
		if sumDays > n {
			left = mid + 1
		} else {
			right = mid
		}
	}
	fmt.Println(right)
}
