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
	count := 0
	for _, c := range str {
		if c >= 'A' && c <= 'Z' {
			count++
		}
	}
	fmt.Println(count)
}
