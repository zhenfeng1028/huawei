package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var ip string
	fmt.Scan(&ip)
	segs := strings.Split(ip, ".")
	if len(segs) != 4 {
		fmt.Println("NO")
		return
	}
	for _, seg := range segs {
		n, err := strconv.ParseInt(seg, 10, 64)
		if err != nil {
			fmt.Println("NO")
			return
		}
		if len(seg) != len(strconv.Itoa(int(n))) {
			fmt.Println("NO")
			return
		}
		if n < 0 || n > 255 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
