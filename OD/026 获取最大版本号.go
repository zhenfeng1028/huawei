package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var v1, v2 string
	fmt.Scan(&v1, &v2)
	strs1 := strings.Split(v1, "-")
	strs2 := strings.Split(v2, "-")
	frontPart1, frontPart2 := strs1[0], strs2[0]
	var milestone1, milestone2 string
	if len(strs1) > 1 {
		milestone1 = strs1[1]
	}
	if len(strs2) > 1 {
		milestone2 = strs2[1]
	}
	ret := compareFrontPart(frontPart1, frontPart2)
	if ret == 1 {
		fmt.Println(v1)
		return
	} else if ret == 2 {
		fmt.Println(v2)
		return
	} else {
		if len(milestone1) != 0 && len(milestone2) != 0 {
			if milestone1 < milestone2 {
				fmt.Println(v2)
				return
			} else { // 1.milestone1 > milestone2 2.两者相等
				fmt.Println(v1)
				return
			}
		} else if len(milestone1) == 0 && len(milestone2) != 0 {
			fmt.Println(v2)
			return
		} else { // 1.milestone1存在，milestone2不存在 2.两个都不存在
			fmt.Println(v1)
			return
		}
	}
}

// s1大返回1，s2大返回2，一样大返回0
func compareFrontPart(s1, s2 string) int {
	vs1 := strings.Split(s1, ".")
	vs2 := strings.Split(s2, ".")
	vs11 := []int{}
	for _, v := range vs1 {
		n, _ := strconv.Atoi(v)
		vs11 = append(vs11, n)
	}
	vs21 := []int{}
	for _, v := range vs2 {
		n, _ := strconv.Atoi(v)
		vs21 = append(vs21, n)
	}
	if vs11[0] > vs21[0] {
		return 1
	} else if vs11[0] < vs21[0] {
		return 2
	}
	if vs11[1] > vs21[1] {
		return 1
	} else if vs11[1] < vs21[1] {
		return 2
	}
	if len(vs11) == 2 && len(vs21) == 2 {
		return 0
	}
	if len(vs11) == 3 && len(vs21) == 3 {
		if vs11[2] > vs21[2] {
			return 1
		} else if vs11[2] < vs21[2] {
			return 2
		} else {
			return 0
		}
	}
	if len(vs11) == 3 {
		return 1
	}
	return 2
}
