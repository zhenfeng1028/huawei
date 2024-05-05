package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str1 := input.Text()
	input.Scan()
	str2 := input.Text()
	fmt.Println(encrypt(str1))
	fmt.Println(decrypt(str2))
}

func encrypt(str string) string {
	res := []rune{}
	for _, c := range str {
		if c >= '0' && c < '9' {
			res = append(res, c+1)
			continue
		}
		if c == '9' {
			res = append(res, '0')
			continue
		}
		if c == 'z' {
			res = append(res, 'A')
			continue
		}
		if c == 'Z' {
			res = append(res, 'a')
			continue
		}
		if unicode.IsLower(c) {
			res = append(res, unicode.ToUpper(c)+1)
		} else {
			res = append(res, unicode.ToLower(c)+1)
		}
	}
	return string(res)
}

func decrypt(str string) string {
	res := []rune{}
	for _, c := range str {
		if c > '0' && c <= '9' {
			res = append(res, c-1)
			continue
		}
		if c == '0' {
			res = append(res, '9')
			continue
		}
		if c == 'a' {
			res = append(res, 'Z')
			continue
		}
		if c == 'A' {
			res = append(res, 'z')
			continue
		}
		if unicode.IsLower(c) {
			res = append(res, unicode.ToUpper(c)-1)
		} else {
			res = append(res, unicode.ToLower(c)-1)
		}
	}
	return string(res)
}
