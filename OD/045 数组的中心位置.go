package main

import "fmt"

func main() {
	var arr []int
	for {
		var tmp int
		n, _ := fmt.Scan(&tmp)
		if n == 0 {
			break
		}
		arr = append(arr, tmp)
	}
	res := -1
	for i := 0; i < len(arr); i++ {
		leftProduct, rightProduct := 1, 1
		if i == 0 {
			leftProduct = 1
			for j := 1; j < len(arr); j++ {
				rightProduct *= arr[j]
			}
		} else if i == len(arr)-1 {
			rightProduct = 1
			for j := len(arr) - 2; j >= 0; j-- {
				leftProduct *= arr[j]
			}
		} else {
			left, right := i-1, i+1
			for left >= 0 {
				leftProduct *= arr[left]
				left--
			}
			for right < len(arr) {
				rightProduct *= arr[right]
				right++
			}
		}
		if leftProduct == rightProduct {
			res = i
			break
		}
	}
	fmt.Println(res)
}
