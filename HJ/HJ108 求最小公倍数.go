package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(a * b / gcd(a, b))
}

// 最大公约数
func gcd(a, b int) int {
	var c int
	for b != 0 {
		c = a % b
		a = b
		b = c
	}
	return a
}
