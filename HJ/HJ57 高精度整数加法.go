package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	m, n := len(str1)-1, len(str2)-1
	carry := 0
	tmp := []int{}
	for m >= 0 && n >= 0 {
		sum := int(str1[m]+str2[n]) - 96 + carry
		carry = 0
		if sum >= 10 {
			carry = 1
			tmp = append(tmp, sum%10)
		} else {
			tmp = append(tmp, sum)
		}
		m--
		n--
	}
	if m < 0 {
		for n >= 0 {
			sum := int(str2[n]) - 48 + carry
			carry = 0
			if sum >= 10 {
				carry = 1
				tmp = append(tmp, sum%10)
			} else {
				tmp = append(tmp, sum)
			}
			n--
		}
	}
	if n < 0 {
		for m >= 0 {
			sum := int(str1[m]) - 48 + carry
			carry = 0
			if sum >= 10 {
				carry = 1
				tmp = append(tmp, sum%10)
			} else {
				tmp = append(tmp, sum)
			}
			m--
		}
	}
	if carry > 0 {
		tmp = append(tmp, carry)
	}
	res := ""
	for i := len(tmp) - 1; i >= 0; i-- {
		res += strconv.Itoa(tmp[i])
	}
	fmt.Println(res)
}
