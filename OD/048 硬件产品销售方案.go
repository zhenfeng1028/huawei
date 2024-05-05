package main

import (
	"fmt"
	"strconv"
)

var amout int
var price []int
var res [][]int

func main() {
	fmt.Scan(&amout)
	var str string
	fmt.Scan(&str)
	tmp := ""
	for _, c := range str {
		if '0' <= c && c <= '9' {
			tmp += string(c)
		} else if c == ',' || c == ']' {
			if len(tmp) > 0 {
				n, _ := strconv.Atoi(tmp)
				price = append(price, n)
				tmp = ""
			}
		}
	}
	var s = []int{}
	dfs(amout, s)
	// 需要对res去重
	fmt.Println(res)
}

func dfs(target int, s []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		tmp := []int{}
		for _, v := range s {
			tmp = append(tmp, v)
		}
		res = append(res, tmp) // 不能直接将s append进去
		return
	}
	for i := 0; i < len(price); i++ {
		if price[i] <= target {
			s = append(s, price[i])
			dfs(target-price[i], s)
			s = s[:len(s)-1]
		}
	}
}
