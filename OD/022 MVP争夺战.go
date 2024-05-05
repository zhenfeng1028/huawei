package main

import "fmt"

var t int        // 分钟数
var scores []int // 分钟得分
var highest int  // 最高得分
var sum int      // 总得分

func main() {
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var score int
		fmt.Scan(&score)
		scores = append(scores, score)
	}
	for _, score := range scores {
		sum += score
		if score > highest {
			highest = score
		}
	}
	for i := highest; i <= sum; i++ {
		if sum%i == 0 {
			p := sum / i
			bucket := make([]int, p)
			if dfs(bucket, 0, i) {
				fmt.Println(i)
				break
			}
		}
	}
}

func dfs(bucket []int, index int, average int) bool {
	if index == len(scores) { // 当所有元素都用完的时候，bucket中的值一定都是average
		return true
	}
	for i := 0; i < len(bucket); i++ {
		if bucket[i]+scores[index] <= average {
			bucket[i] += scores[index]
			if dfs(bucket, index+1, average) {
				return true
			}
			bucket[i] -= scores[index]
		}
	}
	return false
}
