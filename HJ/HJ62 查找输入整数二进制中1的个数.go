package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := []int{}
	for {
		var num int
		n, _ := fmt.Scan(&num)
		if n == 0 {
			break
		}
		binStr := strconv.FormatInt(int64(num), 2)
		count := 0
		for _, c := range binStr {
			if c == '1' {
				count++
			}
		}
		res = append(res, count)
	}
	for _, v := range res {
		fmt.Println(v)
	}
}
