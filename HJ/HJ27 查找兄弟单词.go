package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	strs := strings.Fields(input.Text())
	n := len(strs)
	origin := strs[n-2]
	originChars := make(map[rune]int)
	for _, c := range origin {
		if _, ok := originChars[c]; !ok {
			originChars[c] = 1
		} else {
			originChars[c]++
		}
	}
	candidates := strs[1 : n-2]
	brothers := []string{}
	for _, candidata := range candidates {
		if candidata == origin {
			continue
		}
		if len(candidata) != len(origin) {
			continue
		}
		candiChars := make(map[rune]int)
		for _, c := range candidata {
			if _, ok := candiChars[c]; !ok {
				candiChars[c] = 1
			} else {
				candiChars[c]++
			}
		}
		valid := true
		for c, v := range candiChars {
			if count, ok := originChars[c]; !ok {
				valid = false
				break
			} else {
				if v != count {
					valid = false
					break
				}
			}
		}
		if valid {
			brothers = append(brothers, candidata)
		}
	}
	sort.Strings(brothers)
	fmt.Println(len(brothers))
	k, _ := strconv.Atoi(strs[n-1])
	if k <= len(brothers) {
		fmt.Println(brothers[k-1])
	}
}
