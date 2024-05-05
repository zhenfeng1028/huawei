package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var colors []int // 车辆颜色信息
var window int   // 窗口大小

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	strs := strings.Split(input.Text(), " ")
	for _, str := range strs {
		color, _ := strconv.Atoi(str)
		colors = append(colors, color)
	}
	fmt.Scan(&window)
	windows := []int{}
	for i := 0; i <= len(colors)-window; i++ {
		var (
			zero int
			one  int
			two  int
		)
		for j := i; j < i+window; j++ {
			if colors[j] == 0 {
				zero++
			} else if colors[j] == 1 {
				one++
			} else {
				two++
			}
		}
		windows = append(windows, max(zero, one, two))
	}
	maxCount := 0
	for _, v := range windows {
		if v > maxCount {
			maxCount = v
		}
	}
	fmt.Println(maxCount)
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}
