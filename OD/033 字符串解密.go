package main

import (
	"fmt"
	"sort"
)

type SubStr struct {
	Origin    string `json:"origin"`
	DiffCount int    `json:"diff_count"` // 不同字符数量
}

func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	hashmap := make(map[rune]int)
	for _, c := range str2 {
		if _, ok := hashmap[c]; !ok {
			hashmap[c] = 1
		} else {
			hashmap[c]++
		}
	}
	scrambling := false
	selection := []string{}
	tmp := ""
	for _, c := range str1 {
		if 'g' <= c && c <= 'z' {
			if !scrambling {
				tmp += string(c)
			} else {
				tmp = string(c)
			}
			scrambling = false
		} else {
			if tmp != "" {
				selection = append(selection, tmp)
				tmp = ""
			}
			scrambling = true
		}
	}
	lastChar := str1[len(str1)-1]
	if 'g' <= lastChar && lastChar <= 'z' {
		selection = append(selection, tmp)
	}
	subStrs := make([]*SubStr, 0)
	for _, v := range selection {
		subStr := &SubStr{Origin: v}
		hashmap2 := make(map[rune]int)
		for _, c := range v {
			if _, ok := hashmap2[c]; !ok {
				hashmap2[c] = 1
			} else {
				hashmap2[c]++
			}
		}
		subStr.DiffCount = len(hashmap2)
		subStrs = append(subStrs, subStr)
	}
	sort.Sort(SubStrSlice(subStrs))
	if subStrs[0].DiffCount > len(hashmap) {
		fmt.Println("Not Found")
		return
	}
	count := 0
	for i := 0; i < len(subStrs); i++ {
		if subStrs[i].DiffCount > len(hashmap) && i > 0 {
			count = subStrs[i-1].DiffCount
			break
		}
	}
	if count == 0 {
		count = subStrs[len(subStrs)-1].DiffCount
	}
	strs := []string{}
	for _, v := range subStrs {
		if v.DiffCount == count {
			strs = append(strs, v.Origin)
		}
	}
	if len(strs) == 0 {
		fmt.Println("Not Found")
		return
	}
	sort.Strings(strs)
	fmt.Println(strs[len(strs)-1])
}

type SubStrSlice []*SubStr

func (s SubStrSlice) Len() int      { return len(s) }
func (s SubStrSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SubStrSlice) Less(i, j int) bool {
	return s[i].DiffCount < s[j].DiffCount
}
