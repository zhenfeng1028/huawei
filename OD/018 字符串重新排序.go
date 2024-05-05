package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Word struct {
	Name  string
	Count int
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	words := strings.Split(input.Text(), " ")
	processed := []string{}
	for _, word := range words {
		chars := []rune{}
		for _, c := range word {
			chars = append(chars, c)
		}
		sort.SliceStable(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		processed = append(processed, string(chars))
	}
	hashmap := make(map[string]*Word)
	for _, v := range processed {
		if _, ok := hashmap[v]; !ok {
			hashmap[v] = &Word{
				Name: v,
			}
		}
		hashmap[v].Count++
	}
	tmp := []*Word{}
	for _, v := range hashmap {
		tmp = append(tmp, v)
	}
	sort.Sort(WordSlice(tmp))
	for _, v := range tmp {
		for i := 0; i < v.Count; i++ {
			fmt.Printf("%s ", v.Name)
		}
	}
	fmt.Println()
}

type WordSlice []*Word

func (s WordSlice) Len() int      { return len(s) }
func (s WordSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s WordSlice) Less(i, j int) bool {
	if s[i].Count != s[j].Count {
		return s[i].Count > s[j].Count
	} else if len(s[i].Name) != len(s[j].Name) {
		return len(s[i].Name) < len(s[j].Name)
	} else {
		index := 0
		for k := 0; k < len(s[i].Name); k++ {
			if s[i].Name[k] != s[j].Name[k] {
				index = k
			}
		}
		return s[i].Name[index] < s[j].Name[index]
	}
}
