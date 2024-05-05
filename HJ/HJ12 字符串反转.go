package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	str := ""
	for i := len(input) - 1; i >= 0; i-- {
		str += string(input[i])
	}
	fmt.Println(str)
}
