package main

import (
	"fmt"
	"math"
)

func main() {
	var num float32
	fmt.Scan(&num)
	fmt.Println(math.Floor(float64(num) + 0.5))
}
