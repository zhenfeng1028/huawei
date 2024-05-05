package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	month, _ := strconv.Atoi(scanner.Text())
	fmt.Println(rabbitNum(month))
}

// 迭代
func rabbitNum(n int) int {
	k1, k2, k3 := 1, 0, 0 // k1表示萌新，k2表示不可生育，k3表示可生育
	for i := 2; i <= n; i++ {
		k3 = k3 + k2
		k2 = k1
		k1 = k3
	}
	return k1 + k2 + k3
}

// 动态规划
// 设第n个月的兔子数量为dp[n],第n个月有第n-1个月已有的兔子，同时，可能会有新出生的兔子。
// 由题目可知，每只兔子在第三个月都会生一只兔子，所以第n个月出生的兔子数量等于第n-2月时的兔子数量，
// 即dp[n]=dp[n-1]+dp[n-2]。
func rabbitNum2(n int) int {
	dp := [32]int{}
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
