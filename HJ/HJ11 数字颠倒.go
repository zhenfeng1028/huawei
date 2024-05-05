package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scan(&num)
	if num == 0 {
		fmt.Println(strconv.Itoa(0))
	}
	str := ""
	for num > 0 {
		str += strconv.Itoa(num % 10)
		num /= 10
	}
	fmt.Println(str)
}
