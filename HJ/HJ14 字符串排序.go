package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var count int
	fmt.Scan(&count)
	strs := []string{}
	input := bufio.NewScanner(os.Stdin)
	for i := 0; i < count; i++ {
		input.Scan()
		str := input.Text()
		strs = append(strs, str)
	}
	sort.Strings(strs)
	for _, str := range strs {
		fmt.Println(str)
	}
}
