package main

import "fmt"

var hashmap = map[string]int{
	"A":  1,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
}

var hashmap2 = map[int]string{
	1:  "A",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "10",
	11: "J",
	12: "Q",
	13: "K",
}

var arr [4]int
var visited [4]bool  // 用于判断对应序号的数字有没有被使用
var res []int        // 存放可行的顺序
var symbols []string // 存放数字间的符号

func main() {
	var strs [4]string
	for i := 0; i < 4; i++ {
		fmt.Scan(&strs[i])
	}
	for i := 0; i < len(strs); i++ {
		if v, ok := hashmap[strs[i]]; !ok {
			fmt.Println("ERROR")
			return
		} else {
			arr[i] = v
		}
	}
	if canGet24(0, 0) {
		fmt.Printf("%s%s%s%s%s%s%s\n", hashmap2[res[0]], symbols[2], hashmap2[res[1]], symbols[1], hashmap2[res[2]], symbols[0], hashmap2[res[3]])
	} else {
		fmt.Println("NONE")
	}
}

// 使用递归，used表示当前用了多少个数字，tmp表示对已经使用过的数字进行运算后得到的临时结果
func canGet24(used int, tmp int) bool {
	if used == 4 && tmp == 24 {
		return true
	}
	if used == 0 {
		for i := 0; i < 4; i++ {
			visited[i] = true
			res = append(res, arr[i])
			if canGet24(1, arr[i]) {
				return true
			}
			res = res[:len(res)-1]
			visited[i] = false // 每一轮循环要把访问记录恢复
		}
	} else { // used == 1 || used == 2 || used == 3
		for i := 0; i < 4; i++ {
			if !visited[i] {
				visited[i] = true
				res = append(res, arr[i])
				if canGet24(used+1, tmp+arr[i]) {
					symbols = append(symbols, "+")
					return true
				} else if canGet24(used+1, tmp-arr[i]) {
					symbols = append(symbols, "-")
					return true
				} else if canGet24(used+1, tmp*arr[i]) {
					symbols = append(symbols, "*")
					return true
				} else if canGet24(used+1, tmp/arr[i]) {
					symbols = append(symbols, "/")
					return true
				}
				res = res[:len(res)-1]
				visited[i] = false
			}
		}
		return false
	}
	return false // 不是1～4返回false
}
