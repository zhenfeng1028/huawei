package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type NewsPaper struct {
	Origin string
	Chars  map[rune]int
	Used   bool
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	newsPapers := strings.Split(scanner.Text(), " ")
	scanner.Scan()
	anonymousLetters := strings.Split(scanner.Text(), " ")
	papers := []*NewsPaper{}
	for _, np := range newsPapers {
		paper := &NewsPaper{Origin: np, Chars: make(map[rune]int)}
		for _, c := range np {
			if _, ok := paper.Chars[c]; !ok {
				paper.Chars[c] = 1
			} else {
				paper.Chars[c]++
			}
		}
		papers = append(papers, paper)
	}
	valid := false
	for i := 0; i < len(anonymousLetters); i++ {
		hashmap := map[rune]int{}
		for _, c := range anonymousLetters[i] {
			if _, ok := hashmap[c]; !ok {
				hashmap[c] = 1
			} else {
				hashmap[c]++
			}
		}
		found := true
		for _, paper := range papers {
			found = true
			if paper.Used {
				continue
			}
			if len(paper.Origin) == len(anonymousLetters[i]) {
				for k, v := range hashmap {
					if vv, ok := paper.Chars[k]; !ok {
						found = false
						break
					} else {
						if v != vv {
							found = false
							break
						}
					}
				}
			} else {
				found = false
			}
			if found {
				paper.Used = true
				break
			}
		}
		if !found {
			fmt.Println(false)
			break
		} else {
			if i == len(anonymousLetters)-1 {
				valid = true
			}
		}
	}
	if valid {
		fmt.Println(true)
	}
}
