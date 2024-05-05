package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	longestDis := -1
	for i := 0; i < n; i++ {
		left, right := i, n-1
		for left < right && arr[left] != arr[right] {
			right--
		}
		if left != right {
			if right-left > longestDis {
				longestDis = right - left
			}
		}
	}
	fmt.Println(longestDis)
}
