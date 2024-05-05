package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)
	ss := []string{}
	for _, c := range input {
		ss = append(ss, string(c))
	}
	sort.Strings(ss)
	fmt.Println(strings.Join(ss, ""))
}
