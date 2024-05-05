package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	res := []byte{}
	for i := 0; i < len(input); i++ {
		if '0' <= input[i] && input[i] <= '9' {
			if i == 0 || input[i-1] < '0' || input[i-1] > '9' {
				res = append(res, '*')
			}
		} else {
			if i > 0 && '0' <= input[i-1] && input[i-1] <= '9' {
				res = append(res, '*')
			}
		}
		res = append(res, input[i])
	}
	if '0' <= input[len(input)-1] && input[len(input)-1] <= '9' {
		res = append(res, '*')
	}
	fmt.Println(string(res))
}
