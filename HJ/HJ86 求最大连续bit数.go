package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	binStr := strconv.FormatInt(int64(n), 2)
	maxLen := 0
	for _, v := range strings.Split(binStr, "0") {
		if len(v) > maxLen {
			maxLen = len(v)
		}
	}
	fmt.Println(maxLen)
}
