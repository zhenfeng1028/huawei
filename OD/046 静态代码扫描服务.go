package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache int    // 缓存一个报告的金币数
var ids []int    // 文件标识序列
var scales []int // 文件大小序列
var res int      // 最小金币数

func main() {
	fmt.Scan(&cache)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strs1 := strings.Split(scanner.Text(), " ")
	for _, str := range strs1 {
		n, _ := strconv.Atoi(str)
		ids = append(ids, n)
	}
	scanner.Scan()
	strs2 := strings.Split(scanner.Text(), " ")
	for _, str := range strs2 {
		n, _ := strconv.Atoi(str)
		scales = append(scales, n)
	}
	for i := 0; i < len(ids); i++ {
		if ids[i] != -1 {
			cnt := 1
			for j := i + 1; j < len(ids); j++ {
				if ids[j] == ids[i] {
					cnt++
					ids[j] = -1
				}
			}
			res += min(cache+scales[i], cnt*scales[i]) // 使用缓存和不使用缓存两个中的较小值
		}
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
