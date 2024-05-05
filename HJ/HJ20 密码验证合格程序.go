package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	res := []string{}
	for scanner.Scan() {
		str := scanner.Text()
		if len(str) <= 8 {
			res = append(res, "NG")
			continue
		}
		cmap := make(map[byte][]int)
		lowercase, uppercase, number, other := 0, 0, 0, 0
		for i, c := range str {
			if _, ok := cmap[byte(c)]; !ok {
				cmap[byte(c)] = make([]int, 0)
			}
			cmap[byte(c)] = append(cmap[byte(c)], i)
			if c >= 'a' && c <= 'z' {
				if lowercase == 0 {
					lowercase++
				}
			} else if c >= 'A' && c <= 'Z' {
				if uppercase == 0 {
					uppercase++
				}
			} else if c >= '0' && c <= '9' {
				if number == 0 {
					number++
				}
			} else if other == 0 {
				other++
			}
		}
		if lowercase+uppercase+number+other < 3 {
			res = append(res, "NG")
			continue
		}
		repeat := false
		for _, v := range cmap {
			if len(v) > 1 {
				// preIndex := -1
				subStr := make(map[string]struct{})
				for _, index := range v {
					// if preIndex != -1 {
					// 	if index-preIndex < 3 {	// 子串可以重叠的话就不需要这个判断
					// 		continue
					// 	}
					// }
					if index+3 <= len(str) {
						if _, ok := subStr[str[index:index+3]]; !ok {
							subStr[str[index:index+3]] = struct{}{}
						} else {
							res = append(res, "NG")
							repeat = true
							break
						}
					}
				}
			}
			if repeat {
				break
			}
		}
		if !repeat {
			res = append(res, "OK")
		}
	}
	for _, v := range res {
		fmt.Println(v)
	}
}
