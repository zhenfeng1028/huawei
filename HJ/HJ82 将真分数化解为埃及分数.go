package main

import "fmt"

func main() {
	var a, b int
	for {
		fmt.Scanf("%d/%d", &a, &b)
		for a != 1 {
			if b%a == 0 {
				b /= a
				a = 1
				break
			}
			q := b / a
			fmt.Printf("%d/%d+", 1, 1+q)
			a -= b % a
			b *= (1 + q)
		}
		fmt.Printf("%d/%d\n", a, b)
	}
}
