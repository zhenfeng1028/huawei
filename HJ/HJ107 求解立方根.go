package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Scan(&num)
	sign := -1.0
	if num > 0 {
		sign = 1.0
	} else {
		num = -num
	}
	left, right, res := 0.0, num, 0.0
	if num < 1 {
		right = 1
	} else {
		left = 1
	}
	for {
		mid := left + (right-left)/2
		if math.Abs(mid*mid*mid-num) <= 0.01 {
			res = mid
			break
		} else if mid*mid*mid <= num {
			left = mid
		} else {
			right = mid
		}
	}
	fmt.Printf("%.1f\n", sign*res)
}
