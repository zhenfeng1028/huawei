package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var hashmap = map[rune]string{
	'0': "0",
	'1': "8",
	'2': "4",
	'3': "C",
	'4': "2",
	'5': "A",
	'6': "6",
	'7': "E",
	'8': "1",
	'9': "9",
	'a': "5",
	'b': "D",
	'c': "3",
	'd': "B",
	'e': "7",
	'f': "F",
	'A': "5",
	'B': "D",
	'C': "3",
	'D': "B",
	'E': "7",
	'F': "F",
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	strs := strings.Fields(input.Text())
	combinedStr := strs[0] + strs[1]
	odds, evens := []rune{}, []rune{}
	for i, c := range combinedStr {
		if i%2 == 0 {
			evens = append(evens, c)
		} else {
			odds = append(odds, c)
		}
	}
	sort.SliceStable(odds, func(i, j int) bool {
		return odds[i] < odds[j]
	})
	sort.SliceStable(evens, func(i, j int) bool {
		return evens[i] < evens[j]
	})
	rs := make([]rune, len(odds)+len(evens))
	for i, v := range evens {
		rs[i*2] = v
	}
	for i, v := range odds {
		rs[i*2+1] = v
	}
	res := ""
	for _, r := range string(rs) {
		if v, ok := hashmap[r]; ok {
			res += string(v)
		} else {
			res += string(r)
		}
	}
	fmt.Println(string(res))
}
