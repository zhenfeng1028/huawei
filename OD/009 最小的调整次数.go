package main

import (
	"bufio"
	"fmt"
	"os"
)

var queueNum int // 队列中元素的数量
var flag bool    // 是否需要调整次序
var num int      // 调整的次数

func main() {
	var n int
	fmt.Scan(&n)
	ops := []string{}
	input := bufio.NewScanner(os.Stdin)
	for i := 0; i < 2*n; i++ {
		input.Scan()
		op := input.Text()
		ops = append(ops, op)
	}
	for _, op := range ops {
		if len(op) != 6 { // 非remove操作
			if op[0] == 'h' && queueNum != 0 {
				flag = true // 从队头插入且此时队列中存在元素，下次移出的时候需要调整
			}
			queueNum++
		} else { // remove操作
			if flag {
				num++ // 调整次数加一
			}
			queueNum--
			flag = false // 移出元素之后需要重置flag
		}
	}
	fmt.Println(num)
}
