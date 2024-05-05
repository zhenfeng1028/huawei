package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Scan(&num)
	if num == 1 {
		fmt.Println(1)
	}
	for num != 1 {
		// 加一个质数判断
		if isPrime(num) {
			fmt.Print(num, " ")
			break
		}
		// 最坏情况下需要遍历2~num之间所有数，时间复杂度过高
		for i := 2; i <= num; i++ {
			if num%i == 0 {
				fmt.Print(i, " ")
				num /= i
				break
			}
		}
	}
}

func isPrime(num int) bool {
	flag := true
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			flag = false
			break
		}
	}
	return flag
}
