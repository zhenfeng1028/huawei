package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	strs := strings.Fields(input.Text())
	logs := []int{}
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		logs = append(logs, num)
	}
	firstReportScore := 0
	totalNum := 0
	curNum := 0
	report := false
	for i := 0; i < len(logs); i++ {
		if totalNum+logs[i] >= 100 {
			// 上报
			curNum = 100 - totalNum
			report = true
		} else {
			curNum = logs[i]
		}
		totalNum += curNum
		score := totalNum
		if i > 0 {
			for j := i - 1; j >= 0; j-- {
				score -= (i - j) * logs[j]
			}
		}
		if score > firstReportScore {
			firstReportScore = score
		}
		if report {
			break
		}
	}
	fmt.Println(firstReportScore)
}
