package main

import (
	"fmt"
	"math"
)

var num int     // 商品数量
var days int    // 售货天数
var item []int  // 每件商品最大持有量
var arr [][]int // 各个商品每日的价格

func main() {
	fmt.Scan(&num, &days)
	item = make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scan(&item[i])
	}
	arr = make([][]int, num)
	for i := 0; i < num; i++ {
		arr[i] = make([]int, days)
	}
	for i := 0; i < num; i++ {
		for j := 0; j < days; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	profit := 0
	for i := 0; i < num; i++ {
		profit += getProfit(i) // 各个商品之间的利润是独立的，互不影响
	}
	fmt.Println(profit)
}

func getProfit(index int) int {
	prices := arr[index]
	// 找出售价最小的那天，如果不是最后一天，那么该商品就有利润
	minPrice := math.MaxInt64
	minPriceDay := -1
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			minPriceDay = i
		}
	}
	if minPriceDay != len(prices)-1 {
		// 找出minPriceDay之后最大利润那天
		maxPrice := math.MinInt64
		maxPriceDay := -1
		for i := minPriceDay + 1; i < len(prices); i++ {
			if prices[i] > maxPrice {
				maxPrice = prices[i]
				maxPriceDay = i
			}
		}
		profit := (prices[maxPriceDay] - prices[minPriceDay]) * item[index]
		return profit
	}
	return 0
}
