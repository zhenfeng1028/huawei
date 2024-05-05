package main

import "fmt"

var arr [4]int
var visited [4]bool // 用于判断对应序号的整数有没有被使用

func main() {
	for i := 0; i < 4; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(canGet24(0, 0))
}

// 使用递归，used表示当前用了多少个数字，tmp表示对已经使用过的数字进行运算后得到的临时结果
func canGet24(used int, tmp int) bool {
	if used == 4 && tmp == 24 {
		return true
	}
	if used == 0 {
		for i := 0; i < 4; i++ {
			visited[i] = true
			if canGet24(1, arr[i]) {
				return true
			}
			visited[i] = false // 每一轮循环要把访问记录恢复
		}
	}
	if used == 2 {
		// 对于已经有两个数字参加运算的情况，要考虑两种可能
		// 1.另外两个数字参加运算后再和当前结果运算
		a, b := 0, 0
		for i := 0; i < 4; i++ {
			if !visited[i] {
				if a == 0 {
					a = arr[i]
				} else {
					b = arr[i]
				}
			}
		}
		for _, v := range getAnyRes(a, b) { // 对所有可能的ab结果进行判断
			for _, vv := range getAnyRes(v, tmp) { // 对所有可能的v和tmp的运算结果进行判断
				if vv == 24 {
					return true
				}
			}
		}
		// 2.当前结果与第三个数字参加运算
		for i := 0; i < 4; i++ {
			if !visited[i] {
				visited[i] = true
				if canGet24(3, tmp+arr[i]) || canGet24(3, tmp-arr[i]) ||
					canGet24(3, arr[i]-tmp) || canGet24(3, tmp*arr[i]) {
					return true
				}
				if arr[i] != 0 && canGet24(3, tmp/arr[i]) {
					return true
				}
				if tmp != 0 && canGet24(3, arr[i]/tmp) {
					return true
				}
				visited[i] = false
			}
		}
		return false
	}
	if used == 1 || used == 3 { // 如果用了一个数或者三个数，正常判断所有情况
		for i := 0; i < 4; i++ {
			if !visited[i] {
				visited[i] = true
				if canGet24(used+1, tmp+arr[i]) || canGet24(used+1, tmp-arr[i]) ||
					canGet24(used+1, arr[i]-tmp) || canGet24(used+1, tmp*arr[i]) {
					return true
				}
				if arr[i] != 0 && canGet24(used+1, tmp/arr[i]) {
					return true
				}
				if tmp != 0 && canGet24(used+1, arr[i]/tmp) {
					return true
				}
				visited[i] = false
			}
		}
	}
	return false // 不是1～4返回false
}

func getAnyRes(a, b int) []int {
	res := []int{}
	res = append(res, a+b)
	res = append(res, a-b)
	res = append(res, b-a)
	res = append(res, a*b)
	if a != 0 {
		res = append(res, b/a)
	}
	if b != 0 {
		res = append(res, a/b)
	}
	return res
}
