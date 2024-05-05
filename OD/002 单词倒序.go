package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	sentence := input.Text()
	res := ""
	beginIdx := 0
	length := 0
	wordBegin := true
	for i, c := range sentence {
		if c == ' ' || c == '.' || c == ',' || c == '?' {
			wordBegin = true
			if length > 0 {
				tmp := reverseWord(sentence[beginIdx : beginIdx+length])
				res += tmp
				beginIdx = 0
				length = 0
			}
			res += string(c)
		} else if wordBegin {
			beginIdx = i
			length++
			wordBegin = false
		} else {
			length++
		}
	}
	fmt.Println(res)
}

func reverseWord(word string) string {
	bytes := []byte{}
	for i := len(word) - 1; i >= 0; i-- {
		bytes = append(bytes, word[i])
	}
	return string(bytes)
}
