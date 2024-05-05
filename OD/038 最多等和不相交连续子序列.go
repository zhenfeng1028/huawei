package main

import "fmt"

var n int
var arr [1010]int
var pre [1010]int
var hashmap = map[int][2]int{}

func main() {
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
	}
	pre[0] = 0
	for i := 1; i <= n; i++ {
		pre[i] = pre[i-1] + arr[i]
	}
	for j := 1; j <= n; j++ {
		for i := j; i >= 1; i-- {
			sum := pre[j] - pre[i-1]
			if v, ok := hashmap[sum]; !ok {
				hashmap[sum] = [2]int{j, 1}
			} else {
				if v[0] < i {
					hashmap[sum] = [2]int{j, v[1] + 1}
				}
			}
		}
	}
	res := 0
	for _, v := range hashmap {
		if v[1] > res {
			res = v[1]
		}
	}
	fmt.Println(res)
}
