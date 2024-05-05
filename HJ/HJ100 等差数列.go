package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	termN := 2 + (n-1)*3
	S := n * (2 + termN) / 2
	fmt.Println(S)
}
