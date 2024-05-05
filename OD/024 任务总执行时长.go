package main

import (
	"fmt"
	"sort"
)

func main() {
	var taskA, taskB, num int
	fmt.Scan(&taskA, &taskB, &num)
	res := []int{}
	for i := 0; i <= num; i++ {
		total := i*taskA + (num-i)*taskB
		res = append(res, total)
	}
	sort.Ints(res)
	fmt.Println(res)
}
