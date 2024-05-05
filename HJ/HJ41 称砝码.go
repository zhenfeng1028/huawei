package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var N int
	var categories []int
	var number []int
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	_, _ = strconv.Atoi(input.Text())
	input.Scan()
	strs1 := strings.Split(input.Text(), " ")
	for _, str := range strs1 {
		category, _ := strconv.Atoi(str)
		categories = append(categories, category)
	}
	input.Scan()
	strs2 := strings.Split(input.Text(), " ")
	for _, str := range strs2 {
		num, _ := strconv.Atoi(str)
		number = append(number, num)
	}
	weights := []int{0}
	for i, category := range categories {
		for j := 1; j <= number[i]; j++ {
			tmp := []int{}
			for _, weight := range weights {
				v := weight + category
				tmp = append(tmp, v)
			}
			weights = append(weights, tmp...)
		}
	}
	hashmap := make(map[int]struct{})
	for _, v := range weights {
		if _, ok := hashmap[v]; !ok {
			hashmap[v] = struct{}{}
		}
	}
	fmt.Println(len(hashmap))
}
