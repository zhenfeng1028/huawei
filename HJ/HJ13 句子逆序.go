package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	input = string(data)
	strs := strings.Fields(input)
	for i := len(strs) - 1; i >= 0; i-- {
		fmt.Print(strs[i], " ")
	}
}
