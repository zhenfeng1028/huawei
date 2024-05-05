package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var key string
	fmt.Scan(&key)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strs := strings.Split(scanner.Text(), " ")
	extracted := make([]string, len(strs))
	for i := 0; i < len(strs); i++ {
		var tmp []rune
		for _, c := range strs[i] {
			if 'A' <= c && c <= 'Z' {
				tmp = append(tmp, c+32)
			}
			if 'a' <= c && c <= 'z' {
				tmp = append(tmp, c)
			}
		}
		sort.SliceStable(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		extracted[i] = string(tmp)
	}
	found := false
	for i, v := range extracted {
		if v == key {
			fmt.Println(i + 1)
			found = true
			break
		}
	}
	if !found {
		fmt.Println(-1)
	}
}
