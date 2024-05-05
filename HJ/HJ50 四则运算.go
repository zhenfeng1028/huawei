package main

import (
	"fmt"
	"strconv"
	"unicode"
)

var nums []int     // 记录数字
var symbols []byte // 记录运算符
var flag bool

// 双栈法
func main() {
	var str string
	fmt.Scan(&str)
	str = "(" + str
	str = str + ")" // 整个运算式添加括号
	for i := 0; i < len(str); i++ {
		if str[i] == '(' || str[i] == '[' || str[i] == '{' {
			symbols = append(symbols, '(')
		} else if str[i] == ')' || str[i] == ']' || str[i] == '}' {
			for symbols[len(symbols)-1] != '(' { // 弹出开始计算直到遇到左括号
				compute()
			}
			symbols = symbols[:len(symbols)-1] // 弹出左括号
		} else if flag { // 运算符
			for priority(symbols[len(symbols)-1], str[i]) { // 比较运算符优先级
				compute()
			}
			symbols = append(symbols, str[i]) // 将当前运算符入栈
			flag = false
		} else { // 数字
			j := i
			if str[i] == '+' || str[i] == '-' {
				i++
			}
			for unicode.IsDigit(rune(str[i])) {
				i++
			}
			num, _ := strconv.Atoi(str[j:i])
			nums = append(nums, num)
			i--
			flag = true // 数字结束，下一次就是运算符了
		}
	}
	fmt.Println(nums[0])
}

func compute() {
	a := nums[len(nums)-2]
	b := nums[len(nums)-1]
	nums = nums[:len(nums)-2]
	op := symbols[len(symbols)-1]
	symbols = symbols[:len(symbols)-1]
	res := 0
	switch op {
	case '+':
		res = a + b
	case '-':
		res = a - b
	case '*':
		res = a * b
	case '/':
		res = a / b
	}
	nums = append(nums, res)
}

func priority(symbol1, symbol2 byte) bool {
	if symbol1 == '(' {
		return false
	} else if (symbol1 == '+' || symbol1 == '-') && (symbol2 == '*' || symbol2 == '/') {
		return false
	}
	return true
}
