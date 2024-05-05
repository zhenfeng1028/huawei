package main

import "fmt"

var F []int // 工位序列

func main() {
	for {
		var f int
		n, _ := fmt.Scan(&f)
		if n == 0 {
			break
		}
		F = append(F, f)
	}
	friendliness := []int{}
	for i := 0; i < len(F); i++ {
		if F[i] == 0 {
			left, right := i-1, i+1
			for left >= 0 && F[left] == 1 {
				left--
			}
			for right < len(F) && F[right] == 1 {
				right++
			}
			tmp := right - left - 2
			friendliness = append(friendliness, tmp)
		}
	}
	fmt.Println(friendliness)
	if len(friendliness) == 0 {
		fmt.Println(0)
	}
	maxFriendliness := 0
	for _, v := range friendliness {
		if v > maxFriendliness {
			maxFriendliness = v
		}
	}
	fmt.Println(maxFriendliness)
}
