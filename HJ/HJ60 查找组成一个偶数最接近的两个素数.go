package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	pairs := [][2]int{}
	for i := 1; i <= n/2; i = i + 1 {
		if isPrime(i) && isPrime(n-i) {
			pairs = append(pairs, [2]int{i, n - i})
		}
	}
	minDiff := n
	minDiffIdx := 0
	for i, pair := range pairs {
		if int(math.Abs(float64(pair[1]-pair[0]))) < minDiff {
			minDiff = int(math.Abs(float64(pair[1] - pair[0])))
			minDiffIdx = i
		}
	}
	fmt.Println(pairs[minDiffIdx][0])
	fmt.Println(pairs[minDiffIdx][1])
}

func isPrime(a int) bool {
	if a < 2 {
		return true
	}
	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
