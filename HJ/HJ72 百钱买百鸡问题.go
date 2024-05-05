package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)
	var x, y, z int
	res := [][3]int{}
	for i := 14; i >= 0; i-- {
		if (100-7*i)%4 == 0 {
			x = i
			y = (100 - 7*x) / 4
			z = 100 - x - y
			res = append(res, [3]int{x, y, z})
		}
	}
	for i := len(res) - 1; i >= 0; i-- {
		for _, v := range res[i] {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
