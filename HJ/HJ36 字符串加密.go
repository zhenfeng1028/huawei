package main

import (
	"bufio"
	"fmt"
	"os"
)

var alphabet = []rune{
	'a', 'b', 'c', 'd', 'e', 'f',
	'g', 'h', 'i', 'j', 'k', 'l',
	'm', 'n', 'o', 'p', 'q', 'r',
	's', 't', 'u', 'v', 'w', 'x',
	'y', 'z'}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	key := input.Text()
	input.Scan()
	origin := input.Text()
	hashmap := make(map[rune]struct{})
	chars := []rune{}
	for _, c := range key {
		if _, ok := hashmap[c]; !ok {
			hashmap[c] = struct{}{}
			chars = append(chars, c)
		}
	}
	for _, letter := range alphabet {
		if _, ok := hashmap[letter]; ok {
			continue
		} else {
			chars = append(chars, letter)
		}
	}
	hashmap2 := make(map[rune]rune)
	for i, c := range chars {
		hashmap2[alphabet[i]] = c
	}
	rs := []rune{}
	for _, c := range origin {
		if _, ok := hashmap2[c]; ok {
			rs = append(rs, hashmap2[c])
		} else {
			rs = append(rs, c)
		}
	}
	fmt.Println(string(rs))
}
