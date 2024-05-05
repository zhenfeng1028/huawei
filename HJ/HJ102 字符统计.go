package main

import (
	"fmt"
	"sort"
)

type Char struct {
	char  byte
	count int
}

func main() {
	var input string
	fmt.Scan(&input)
	hashmap := make(map[byte]int)
	for i := 0; i < len(input); i++ {
		if _, ok := hashmap[input[i]]; !ok {
			hashmap[input[i]] = 1
		} else {
			hashmap[input[i]]++
		}
	}
	chars := []Char{}
	for k, v := range hashmap {
		c := Char{char: k, count: v}
		chars = append(chars, c)
	}
	sort.Sort(CharSlice(chars))
	for _, c := range chars {
		fmt.Printf("%c", c.char)
	}
	fmt.Println()
}

type CharSlice []Char

func (s CharSlice) Len() int      { return len(s) }
func (s CharSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s CharSlice) Less(i, j int) bool {
	if s[i].count != s[j].count {
		return s[i].count > s[j].count
	} else {
		return s[i].char < s[j].char
	}
}
