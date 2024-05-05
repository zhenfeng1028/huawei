package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scan(&input)
	hexadecimal_num := input[2:]
	s, _ := strconv.ParseInt(hexadecimal_num, 16, 32)
	fmt.Printf("%v\n", s)
}
