package main

import (
	"fmt"
	"sort"
)

func main() {
	var n1, n2 int
	var arr1, arr2 []int
	fmt.Scan(&n1)
	for i := 0; i < n1; i++ {
		var tmp int
		fmt.Scan(&tmp)
		arr1 = append(arr1, tmp)
	}
	fmt.Scan(&n2)
	for i := 0; i < n2; i++ {
		var tmp int
		fmt.Scan(&tmp)
		arr2 = append(arr2, tmp)
	}
	hashmap := make(map[int]struct{})
	for _, v := range arr1 {
		if _, ok := hashmap[v]; !ok {
			hashmap[v] = struct{}{}
		}
	}
	for _, v := range arr2 {
		if _, ok := hashmap[v]; !ok {
			hashmap[v] = struct{}{}
		}
	}
	res := []int{}
	for k := range hashmap {
		res = append(res, k)
	}
	sort.Ints(res)
	for _, v := range res {
		fmt.Printf("%d", v)
	}
	fmt.Println()
}
