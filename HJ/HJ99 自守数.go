package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num int
	fmt.Scan(&num)
	count := 0
	for i := 0; i <= num; i++ {
		str1 := strconv.Itoa(i)
		str2 := strconv.Itoa(i * i)
		// if str2[len(str2)-len(str1):] == str1 {
		// 	count++
		// }
		if strings.HasSuffix(str2, str1) {
			count++
		}
	}
	fmt.Println(count)
}
