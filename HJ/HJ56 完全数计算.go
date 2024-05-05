package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	count := 0
	for i := 1; i <= n; i++ {
		divisors := []int{}
		// 找出除了自身所有的约数
		for j := 1; j < i; j++ {
			if i%j == 0 {
				divisors = append(divisors, j)
			}
		}
		sum := 0
		for _, divisor := range divisors {
			sum += divisor
		}
		if sum == i {
			count++
		}
	}
	fmt.Println(count)
}
