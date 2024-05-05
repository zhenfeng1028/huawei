package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	count := 0
	for i := 1; i <= n; i++ {
		if i%7 == 0 {
			count++
		} else {
			j := i
			for j > 0 {
				if j%10 == 7 {
					count++
					break
				}
				j = j / 10
			}
		}
	}
	fmt.Println(count)
}
