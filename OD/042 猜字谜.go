package main

import (
	"fmt"
	"sort"
	"strings"
)

type Candidate struct {
	Origin string
	Sorted string
	Chars  map[rune]int
}

func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	strs1 := strings.Split(str1, ",")
	strs2 := strings.Split(str2, ",")
	cands := []*Candidate{}
	for _, str := range strs2 {
		chars := []rune{}
		cand := &Candidate{
			Origin: str,
			Chars:  make(map[rune]int),
		}
		for _, c := range str {
			chars = append(chars, c)
			if _, ok := cand.Chars[c]; !ok {
				cand.Chars[c] = 1
			}
		}
		sort.SliceStable(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		cand.Sorted = string(chars)
		cands = append(cands, cand)
	}
	res := []string{}
	for _, str := range strs1 {
		chars := []rune{}
		hashmap := make(map[rune]int)
		for _, c := range str {
			chars = append(chars, c)
			if _, ok := hashmap[c]; !ok {
				hashmap[c] = 1
			}
		}
		sort.SliceStable(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		found := true
		for _, cand := range cands {
			found = true
			if cand.Sorted == string(chars) {
				res = append(res, cand.Origin)
				break
			}
			if len(cand.Chars) != len(hashmap) {
				found = false
			} else {
				for k, _ := range hashmap {
					if _, ok := cand.Chars[k]; !ok {
						found = false
						break
					}
				}
			}
			if found {
				res = append(res, cand.Origin)
				break
			}
		}
		if !found {
			res = append(res, "not found")
		}
	}
	fmt.Println(strings.Join(res, ","))
}
