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
	words := []string{}
	tmp := ""
	for _, c := range str {
		if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
			tmp += string(c)
		} else {
			if len(tmp) > 0 {
				words = append(words, tmp)
				tmp = ""
			}
		}
	}
	if len(tmp) > 0 {
		words = append(words, tmp)
	}
	for i := len(words) - 1; i >= 0; i-- {
		fmt.Printf("%s ", words[i])
	}
}
