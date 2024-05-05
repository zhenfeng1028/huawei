package main

import (
	"fmt"
	"strconv"
	"strings"
)

var hashmap = map[int]string{
	1:  "壹",
	2:  "贰",
	3:  "叁",
	4:  "肆",
	5:  "伍",
	6:  "陆",
	7:  "柒",
	8:  "捌",
	9:  "玖",
	10: "拾",
	11: "拾壹",
	12: "拾贰",
	13: "拾叁",
	14: "拾肆",
	15: "拾伍",
	16: "拾陆",
	17: "拾柒",
	18: "拾捌",
	19: "拾玖",
}

var res string

func main() {
	var input string
	fmt.Scan(&input)
	tmptrtmp := strings.Split(input, ".")
	integer, _ := strconv.Atoi(tmptrtmp[0])
	decimal, _ := strconv.Atoi(tmptrtmp[1])
	if integer > 0 {
		res = "人民币" + f(integer) + "元"
	} else {
		res = "人民币"
	}
	if decimal == 0 {
		res += "整"
	} else if decimal < 10 {
		res += hashmap[decimal] + "分"
	} else if decimal%10 == 0 {
		res += hashmap[decimal/10] + "角"
	} else {
		res += hashmap[decimal/10] + "角" + hashmap[decimal%10] + "分"
	}
	fmt.Println(res)
}

func f(n int) (tmp string) {
	if n < 20 {
		tmp += hashmap[n]
	} else if n < 100 {
		if n%10 >= 1 {
			tmp += f(n/10) + "拾" + f(n%10)
		} else {
			tmp += f(n/10) + "拾"
		}
	} else if n < 1000 {
		if n%100 >= 10 {
			tmp += f(n/100) + "佰" + f(n%100)
		} else if n%100 > 0 {
			tmp += f(n/100) + "佰零" + f(n%100)
		} else {
			tmp += f(n/100) + "佰"
		}
	} else if n < 10000 {
		if n%1000 >= 100 {
			tmp += f(n/1000) + "仟" + f(n%1000)
		} else if n%1000 > 0 {
			tmp += f(n/1000) + "仟零" + f(n%1000)
		} else {
			tmp += f(n/1000) + "仟"
		}
	} else if n < 100000000 {
		if n%10000 >= 1000 {
			tmp += f(n/10000) + "万" + f(n%10000)
		} else if n%10000 > 0 {
			tmp += f(n/10000) + "万零" + f(n%10000)
		} else {
			tmp += f(n/10000) + "万"
		}
	} else {
		if n%100000000 >= 10000000 {
			tmp += f(n/100000000) + "亿" + f(n%100000000)
		} else if n%100000000 > 0 {
			tmp += f(n/100000000) + "亿零" + f(n%100000000)
		} else {
			tmp += f(n/100000000) + "亿"
		}
	}
	return tmp
}
