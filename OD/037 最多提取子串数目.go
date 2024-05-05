package main

import "fmt"

var A, B string
var used [100]bool // 字符串A中的某个字符是否用过

func main() {
	fmt.Scan(&A, &B)
	count := 0
	finished := false
	for {
		next := 0
		for i := 0; i < len(B); i++ {
			found := false
			for j := next; j < len(A); j++ {
				if A[j] == B[i] && !used[j] {
					found = true
					used[j] = true
					next = j + 1
					break
				}
			}
			if !found {
				finished = true
				break
			}
			if i == len(B)-1 && found {
				count++
			}
		}
		if finished {
			break
		}
	}
	fmt.Println(count)
}
