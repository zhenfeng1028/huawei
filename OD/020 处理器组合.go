package main

import (
	"bufio"
	"fmt"
	"os"
)

var available []int    // 可用处理器
var list1, list2 []int // 链路1和链路2
var num int            // 任务申请数量

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	for _, c := range input.Text() {
		if '0' <= c && c <= '9' {
			available = append(available, int(c-'0'))
		}
	}
	fmt.Scan(&num)
	for _, p := range available {
		if p <= 3 {
			list1 = append(list1, p)
		} else {
			list2 = append(list2, p)
		}
	}
	res := [][]int{}
	var hasOne, hasTwo, hasThree, hasFour bool
	if num == 1 {
		if len(list1) == 1 {
			hasOne = true
			res = append(res, list1)
		}
		if len(list2) == 1 {
			hasOne = true
			res = append(res, list2)
		}
		if !hasOne && len(list1) == 3 {
			hasThree = true
			for _, v := range list1 {
				res = append(res, []int{v})
			}
		}
		if !hasOne && len(list2) == 3 {
			hasThree = true
			for _, v := range list2 {
				res = append(res, []int{v})
			}
		}
		if !hasOne && !hasThree && len(list1) == 2 {
			hasTwo = true
			for _, v := range list1 {
				res = append(res, []int{v})
			}
		}
		if !hasOne && !hasThree && len(list2) == 2 {
			hasTwo = true
			for _, v := range list2 {
				res = append(res, []int{v})
			}
		}
		if !hasOne && !hasThree && !hasTwo && len(list1) == 4 {
			for _, v := range list1 {
				res = append(res, []int{v})
			}
		}
		if !hasOne && !hasThree && !hasTwo && len(list2) == 4 {
			for _, v := range list2 {
				res = append(res, []int{v})
			}
		}
	} else if num == 2 {
		if len(list1) == 2 {
			hasTwo = true
			res = append(res, list1)
		}
		if len(list2) == 2 {
			hasTwo = true
			res = append(res, list2)
		}
		if !hasTwo && len(list1) == 4 {
			hasFour = true
			for i := 0; i < 3; i++ {
				for j := i + 1; j < 4; j++ {
					res = append(res, []int{list1[i], list1[j]})
				}
			}
		}
		if !hasTwo && len(list2) == 4 {
			hasFour = true
			for i := 0; i < 3; i++ {
				for j := i + 1; j < 4; j++ {
					res = append(res, []int{list2[i], list2[j]})
				}
			}
		}
		if !hasTwo && !hasFour && len(list1) == 3 {
			for i := 0; i < 2; i++ {
				for j := i + 1; j < 3; j++ {
					res = append(res, []int{list1[i], list1[j]})
				}
			}
		}
		if !hasTwo && !hasFour && len(list2) == 3 {
			for i := 0; i < 2; i++ {
				for j := i + 1; j < 3; j++ {
					res = append(res, []int{list2[i], list2[j]})
				}
			}
		}
	} else if num == 4 {
		if len(list1) == 4 {
			res = append(res, list1)
		}
		if len(list2) == 4 {
			res = append(res, list2)
		}
	} else { // num == 8
		if len(list1) == 4 && len(list2) == 4 {
			tmp := []int{}
			for _, v := range list1 {
				tmp = append(tmp, v)
			}
			for _, v := range list2 {
				tmp = append(tmp, v)
			}
			res = append(res, tmp)
		}
	}
	fmt.Println(res)
}
