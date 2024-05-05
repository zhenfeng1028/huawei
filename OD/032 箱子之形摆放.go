package main

import "fmt"

func main() {
	var str string
	var width int
	fmt.Scan(&str, &width)
	strs := []string{}
	for i := 0; i < len(str)/width; i++ {
		strs = append(strs, str[i*width:(i+1)*width])
	}
	lastIndex := len(str) / width * width
	if len(str)%width != 0 {
		strs = append(strs, str[lastIndex:])
	}
	res := []string{}
	for i := 0; i < width; i++ {
		tmp := ""
		for j := 0; j < len(strs); j++ {
			if j%2 == 0 {
				if j == len(strs)-1 && i >= len(strs[j]) {
					continue
				}
				tmp += string(strs[j][i])
			} else {
				if j == len(strs)-1 && i <= len(strs[j]) {
					continue
				}
				tmp += string(strs[j][width-1-i])
			}
		}
		res = append(res, tmp)
	}
	for _, v := range res {
		fmt.Println(v)
	}
}
