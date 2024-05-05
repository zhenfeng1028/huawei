package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	res := ""
	for _, c := range str {
		if '0' <= c && c <= '9' {
			res += string(c)
		}
		if 'a' <= c && c <= 'c' {
			res += "2"
		} else if 'd' <= c && c <= 'f' {
			res += "3"
		} else if 'g' <= c && c <= 'i' {
			res += "4"
		} else if 'j' <= c && c <= 'l' {
			res += "5"
		} else if 'm' <= c && c <= 'o' {
			res += "6"
		} else if 'p' <= c && c <= 's' {
			res += "7"
		} else if 't' <= c && c <= 'v' {
			res += "8"
		} else if 'w' <= c && c <= 'z' {
			res += "9"
		}
		if 'A' <= c && c <= 'Y' {
			res += string(byte(c) + 33)
		} else if c == 'Z' {
			res += "a"
		}
	}
	fmt.Println(res)
}
