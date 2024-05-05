package main

import "fmt"

func main() {
	var password string
	fmt.Scan(&password)

	var score int
	var exsitLetter bool
	var existUpCase, existLowerCase bool
	var digitCount, symbolCount int

	// 密码长度
	n := len(password)
	switch {
	case n <= 4:
		score += 5
	case 5 <= n && n <= 7:
		score += 10
	case n >= 8:
		score += 25
	}

	for _, c := range password {
		if 'a' <= c && c <= 'z' {
			exsitLetter = true
			existLowerCase = true
		} else if 'A' <= c && c <= 'Z' {
			exsitLetter = true
			existUpCase = true
		} else if '0' <= c && c <= '9' {
			digitCount++
		} else {
			symbolCount++
		}
	}

	// 字母情况
	switch {
	case !existLowerCase && !existUpCase:
		score += 0
	case existLowerCase && existUpCase:
		score += 20
	default:
		score += 10
	}

	// 数字情况
	switch {
	case digitCount == 0:
		score += 0
	case digitCount == 1:
		score += 10
	default:
		score += 20
	}

	// 符号情况
	switch {
	case symbolCount == 0:
		score += 0
	case symbolCount == 1:
		score += 10
	default:
		score += 25
	}

	// 奖励情况
	if existLowerCase && existUpCase && digitCount > 0 && symbolCount > 0 {
		score += 5
	} else if exsitLetter && digitCount > 0 && symbolCount > 0 {
		score += 3
	} else if exsitLetter && digitCount > 0 {
		score += 2
	}

	// 评分
	switch {
	case score >= 90:
		fmt.Println("VERY_SECURE")
	case score >= 80:
		fmt.Println("SECURE")
	case score >= 70:
		fmt.Println("VERY_STRONG")
	case score >= 60:
		fmt.Println("STRONG")
	case score >= 50:
		fmt.Println("AVERAGE")
	case score >= 25:
		fmt.Println("WEAK")
	default:
		fmt.Println("VERY_WEAK")
	}
}
