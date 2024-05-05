package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Score int
}

var ss = []Student{}

func main() {
	var n int
	fmt.Scan(&n)
	var order int
	fmt.Scan(&order)
	for i := 0; i < n; i++ {
		stu := Student{}
		fmt.Scan(&stu.Name, &stu.Score)
		curLen := len(ss)
		for i, s := range ss {
			if order == 0 && stu.Score > s.Score {
				insert(i, stu)
				break

			} else if order == 1 && stu.Score < s.Score {
				insert(i, stu)
				break
			}
		}
		if len(ss) == curLen {
			ss = append(ss, stu)
		}
	}
	for _, v := range ss {
		fmt.Printf("%s %d\n", v.Name, v.Score)
	}
}

func insert(i int, stu Student) {
	tmp1, tmp2 := ss[:i], ss[i:]
	tmp := []Student{}
	tmp = append(tmp, tmp1...)
	tmp = append(tmp, stu)
	tmp = append(tmp, tmp2...)
	ss = tmp
}
