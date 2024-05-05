package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	distance := float64(n)
	riseHeight := 0.0
	for i := 1; i < 5; i++ {
		riseHeight = float64(n) / math.Pow(float64(2), float64(i))
		distance += 2 * riseHeight
	}
	fmt.Printf("%f\n%f\n", distance, riseHeight/2)
}
