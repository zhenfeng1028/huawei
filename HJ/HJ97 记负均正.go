package main

import "fmt"

func main() {
	var n int
	var arr []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		arr = append(arr, tmp)
	}
	minusCount := 0
	zeroCount := 0
	plusSum := 0.0
	for _, v := range arr {
		if v < 0 {
			minusCount++
		} else if v == 0 {
			zeroCount++
		} else {
			plusSum += float64(v)
		}
	}
	average := plusSum / float64(n-minusCount-zeroCount)
	if plusSum == 0 {
		fmt.Printf("%d %.1f \n", minusCount, 0.0)
	} else {
		fmt.Printf("%d %.1f \n", minusCount, average)
	}
}
