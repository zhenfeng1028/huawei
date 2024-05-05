package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	res := []int{}
	for scanner.Scan() {
		empty := scanner.Text()
		n, _ := strconv.Atoi(empty)
		if n == 0 {
			break
		}
		res = append(res, drink(n))
	}
	for _, v := range res {
		fmt.Println(v)
	}
}

func drink(n int) int {
	quotient := n / 3
	mod := n % 3
	if quotient == 0 && mod == 1 {
		return 0
	}
	if quotient == 1 && mod == 0 {
		return 1
	}
	if quotient == 0 && mod == 2 {
		return 1
	}
	return quotient + drink(quotient+mod)
}
