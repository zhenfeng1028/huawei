package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	quotation := 0
	beginIdx := -1
	wordLen := 0
	res := []string{}
	for i := 0; i < len(str); i++ {
		if str[i] == '"' && quotation == 0 {
			beginIdx = i + 1
			quotation = 1
			for j := i + 1; j < len(str); j++ {
				if str[j] == '"' {
					quotation = 0
					res = append(res, str[beginIdx:beginIdx+wordLen])
					i = i + wordLen + 1
					wordLen = 0
					break
				}
				if quotation == 1 {
					wordLen++
				}
			}
		} else if str[i] != ' ' && (i == 0 || i >= 1 && str[i-1] == ' ') {
			beginIdx = i
			wordLen++
		} else if str[i] != ' ' {
			wordLen++
		} else if str[i] == ' ' && (i >= 1 && str[i-1] == '"') {

		} else if str[i] == ' ' && beginIdx != -1 {
			res = append(res, str[beginIdx:beginIdx+wordLen])
			wordLen = 0
		}
	}
	if str[len(str)-1] != ' ' && str[len(str)-1] != '"' { // 把最后一个参数补上
		res = append(res, str[beginIdx:beginIdx+wordLen])
	}
	fmt.Println(len(res))
	for _, v := range res {
		fmt.Println(v)
	}
}
