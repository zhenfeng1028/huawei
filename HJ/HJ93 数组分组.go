package main

import "fmt"

var other []int // 保存既不是3的倍数也不是5的倍数

func main() {
	var n int
	fmt.Scan(&n)
	var arr []int
	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		arr = append(arr, tmp)
	}

	var arr3, arr5 []int
	var sum, sum3 int

	for _, v := range arr {
		if v%5 == 0 {
			arr5 = append(arr5, v)
		} else if v%3 == 0 {
			arr3 = append(arr3, v)
			sum3 += v
		} else {
			other = append(other, v)
		}
		sum += v
	}
	if sum%2 != 0 {
		fmt.Println(false)
		return
	}

	target := sum/2 - sum3 // 在other中寻找是否有元素集合为target
	if target == 0 {
		fmt.Println(true)
		return
	}

	fmt.Println(f(0, target))
}

func f(i int, target int) bool {
	if i >= len(other) {
		return target == 0
	}
	return f(i+1, target-other[i]) || f(i+1, target)
}
