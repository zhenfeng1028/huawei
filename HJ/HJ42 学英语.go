package main

import (
	"fmt"
)

var hashmap = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func main() {
	var number int
	fmt.Scan(&number)
	if number >= 1 && number <= 20 {
		fmt.Println(hashmap[number])
	} else if number > 20 && number < 100 {
		fmt.Println(dozen(number))
	} else if number >= 100 && number < 1000 {
		fmt.Println(hundred(number))
	} else if number >= 1000 && number < 1000000 {
		fmt.Println(thousand(number))
	} else {
		fmt.Println(million(number))
	}
}

func dozen(num int) string {
	n := num / 10
	mod := num % 10
	if mod == 0 {
		return hashmap[num]
	} else {
		return hashmap[n*10] + " " + hashmap[mod]
	}
}

func hundred(num int) string {
	n1 := num / 100
	mod1 := num % 100
	n2 := mod1 - mod1/10
	mod2 := num % 10
	if mod1 == 0 {
		return hashmap[n1] + " hundred"
	} else if mod2 == 0 || mod1 > 10 && mod1 < 20 {
		return hashmap[n1] + " hundred" + " and " + hashmap[mod1]
	} else if n2 > 0 && n2 < 10 {
		return hashmap[n1] + " hundred" + " and " + hashmap[mod2]
	} else {
		return hashmap[n1] + " hundred" + " and " + dozen(mod1)
	}
}

func thousand(num int) string {
	n := num / 1000
	mod := num % 1000
	if mod == 0 {
		if n >= 1 && n <= 20 {
			return hashmap[n] + " thousand"
		} else if n > 20 && n < 100 {
			return dozen(n) + " thousand"
		} else { // n >= 100 && n < 1000
			return hundred(n) + " thousand"
		}
	} else {
		if n >= 1 && n <= 20 {
			if mod >= 1 && mod <= 20 {
				return hashmap[n] + " thousand " + hashmap[mod]
			} else if mod > 20 && mod < 100 {
				return hashmap[n] + " thousand " + dozen(mod)
			} else {
				return hashmap[n] + " thousand " + hundred(mod)
			}
		} else if n > 20 && n < 100 {
			if mod >= 1 && mod <= 20 {
				return dozen(n) + " thousand " + hashmap[mod]
			} else if mod > 20 && mod < 100 {
				return dozen(n) + " thousand " + dozen(mod)
			} else {
				return dozen(n) + " thousand " + hundred(mod)
			}
		} else { // n >= 100 && n < 1000
			if mod >= 1 && mod <= 20 {
				return hundred(n) + " thousand " + hashmap[mod]
			} else if mod > 20 && mod < 100 {
				return hundred(n) + " thousand " + dozen(mod)
			} else {
				return hundred(n) + " thousand " + hundred(mod)
			}
		}
	}
}

func million(num int) string {
	n := num / 1000000
	mod := num % 1000000
	if mod == 0 {
		return hashmap[n] + " million"
	} else {
		return hashmap[n] + " million" + " " + thousand(mod)
	}
}
