package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int64
	fmt.Scan(&num)
	str := strconv.FormatInt(num, 2)
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '1' {
			count++
		}
	}
	fmt.Println(count)
}

// 不使用内置函数转换成二进制
func method2() {
	var num int32
	fmt.Scan(&num)
	count := 0
	for i := 0; i < 32; i++ {
		if (num>>i)&1 == 1 {
			count++
		}
	}
	fmt.Println(count)
}
