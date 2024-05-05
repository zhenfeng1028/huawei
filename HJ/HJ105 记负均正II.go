package main

import "fmt"

func main() {
	var arr []int
	for {
		var tmp int
		n, _ := fmt.Scan(&tmp)
		if n == 0 {
			break
		}
		arr = append(arr, tmp)
	}
	minusCount := 0
	nonNegSum := 0.0
	for _, v := range arr {
		if v < 0 {
			minusCount++
		} else {
			nonNegSum += float64(v)
		}
	}
	if minusCount == len(arr) {
		fmt.Printf("%d\n%.1f\n", minusCount, 0.0)
	} else {
		fmt.Printf("%d\n%.1f\n", minusCount, nonNegSum/float64(len(arr)-minusCount))
	}
}
