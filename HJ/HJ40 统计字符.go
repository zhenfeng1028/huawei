package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	letter, digit, space, other := 0, 0, 0, 0
	for _, c := range str {
		if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
			letter++
		} else if c >= '0' && c <= '9' {
			digit++
		} else if c == ' ' {
			space++
		} else {
			other++
		}
	}
	fmt.Printf("%d\n%d\n%d\n%d\n", letter, space, digit, other)
}
