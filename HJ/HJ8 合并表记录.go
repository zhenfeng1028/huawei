package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	hashmap := make(map[int]int)
	input := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		input.Scan()
		kv := input.Text()
		strs := strings.Split(kv, " ")
		k, _ := strconv.Atoi(strs[0])
		v, _ := strconv.Atoi(strs[1])
		hashmap[k] += v
	}
	ks := []int{}
	for k, _ := range hashmap {
		ks = append(ks, k)
	}
	sort.Ints(ks)
	for _, v := range ks {
		fmt.Printf("%d %d\n", v, hashmap[v])
	}
}
