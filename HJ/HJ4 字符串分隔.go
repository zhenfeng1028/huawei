package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)
	if len(input) == 0 {
		return
	}
	n := len(input) / 8
	mod := len(input) % 8
	for i := 0; i < n; i++ {
		fmt.Println(input[i*8 : i*8+8])
	}
	str := input[8*n:]
	if mod > 0 {
		for i := 0; i < 8-mod; i++ {
			str += "0"
		}
		fmt.Println(str)
	}
}
