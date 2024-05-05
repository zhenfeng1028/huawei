package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	strs := strings.Split(input.Text(), ";")
	row, col := 0, 0
	for _, str := range strs {
		if len(str) <= 1 {
			continue
		}
		n, err := strconv.Atoi(str[1:])
		if err != nil {
			continue
		}
		switch str[0] {
		case 'A':
			row -= n
		case 'D':
			row += n
		case 'W':
			col += n
		case 'S':
			col -= n
		default:
			continue
		}
	}
	fmt.Printf("%d,%d\n", row, col)
}
