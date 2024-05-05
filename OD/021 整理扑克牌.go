package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	Num   int `json:"num"`
	Count int `json:"count"`
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	cards := []int{}
	strs := strings.Split(input.Text(), " ")
	for _, str := range strs {
		card, _ := strconv.Atoi(str)
		cards = append(cards, card)
	}
	hashmap := make(map[int]*Card)
	for _, card := range cards {
		if _, ok := hashmap[card]; !ok {
			hashmap[card] = &Card{
				Num: card,
			}
		}
		hashmap[card].Count++
	}
	cs := []*Card{}
	for _, v := range hashmap {
		cs = append(cs, v)
	}
	sort.Sort(CardSlice(cs))
	res := []int{}
	single := []int{}
	oddTriple := true
	for _, v := range cs {
		if v.Count > 3 || v.Count == 2 {
			for i := 0; i < v.Count; i++ {
				res = append(res, v.Num)
			}
		}
		if oddTriple && v.Count == 3 {
			for i := 0; i < 3; i++ {
				res = append(res, v.Num)
			}
			oddTriple = false
			continue
		}
		if !oddTriple && v.Count == 3 {
			for i := 0; i < 2; i++ {
				res = append(res, v.Num)
			}
			single = append(single, v.Num)
			oddTriple = true
		}
		if v.Count == 1 {
			single = append(single, v.Num)
		}
	}
	sort.Ints(single)
	for i := len(single) - 1; i >= 0; i-- {
		res = append(res, single[i])
	}
	for _, v := range res {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

type CardSlice []*Card

func (s CardSlice) Len() int      { return len(s) }
func (s CardSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s CardSlice) Less(i, j int) bool {
	if s[i].Count != s[j].Count {
		return s[i].Count > s[j].Count
	} else {
		return s[i].Num > s[j].Num
	}
}
