package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	chars := []rune(scanner.Text())
	letters := []rune{}
	otherChars := make([]bool, len(chars))
	for i, c := range chars {
		if unicode.IsLetter(c) {
			letters = append(letters, c)
			continue
		}
		otherChars[i] = true
	}
	sort.SliceStable(letters, func(i, j int) bool {
		ci, cj := letters[i], letters[j]
		return unicode.ToLower(ci) < unicode.ToLower(cj)
	})
	for i, c := range chars {
		if otherChars[i] {
			fmt.Printf("%c", c)
		} else {
			fmt.Printf("%c", letters[0])
			letters = letters[1:]
		}
	}
}
