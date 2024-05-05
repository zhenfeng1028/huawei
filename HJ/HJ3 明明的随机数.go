package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	hashmap := make(map[int]struct{})
	input.Scan()
	N, _ := strconv.Atoi(input.Text())
	for i := 0; i < N; i++ {
		input.Scan()
		num, _ := strconv.Atoi(input.Text())
		hashmap[num] = struct{}{}
	}
	tmp := []int{}
	for k, _ := range hashmap {
		tmp = append(tmp, k)
	}
	sort.Ints(tmp)
	for _, v := range tmp {
		fmt.Println(v)
	}
}
