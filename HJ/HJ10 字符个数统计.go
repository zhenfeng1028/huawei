package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	hashmap := make(map[byte]struct{})
	for i := 0; i < len(input); i++ {
		if _, ok := hashmap[input[i]]; !ok {
			hashmap[input[i]] = struct{}{}
		}
	}
	fmt.Println(len(hashmap))
}
