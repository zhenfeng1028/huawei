package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	for i := len(input) - 1; i >= 0; i-- {
		fmt.Printf("%c", input[i])
	}
	fmt.Println()
}
