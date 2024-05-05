package main

import (
	"fmt"
	"strings"
)

func main() {
	var shorts, longs []string
	for {
		var short, long string
		n, _ := fmt.Scan(&short, &long)
		if n == 0 {
			break
		} else {
			shorts = append(shorts, short)
			longs = append(longs, long)
		}
	}
	res := []bool{}
	for i := 0; i < len(shorts); i++ {
		var contain bool = true
		for _, c := range shorts[i] {
			if !strings.Contains(longs[i], string(c)) {
				res = append(res, false)
				contain = false
				break
			}
		}
		if contain {
			res = append(res, true)
		}
	}
	for _, v := range res {
		fmt.Println(v)
	}
}
