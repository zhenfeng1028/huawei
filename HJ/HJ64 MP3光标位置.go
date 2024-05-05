package main

import "fmt"

func main() {
	var n int
	var ops string
	fmt.Scan(&n, &ops)
	cursor := 1 // 当前光标所在位置
	first := 1  // 当前屏幕显示页的第一首歌曲的编号
	if n <= 4 {
		for _, op := range ops {
			if cursor == 1 && op == 'U' { // 特殊向上翻页
				cursor = n
			} else if cursor == n && op == 'D' { // 特殊向下翻页
				cursor = 1
			} else if op == 'U' { // 一般向上翻页
				cursor--
			} else { // 一般向下翻页
				cursor++
			}
		}
		for i := first; i <= n; i++ {
			fmt.Printf("%d ", i)
		}
		fmt.Println()
		fmt.Println(cursor)
	} else {
		for _, op := range ops {
			if cursor == 1 && op == 'U' { // 特殊向上翻页
				cursor = n
				first = n - 3
			} else if cursor == n && op == 'D' { // 特殊向下翻页
				cursor = 1
				first = 1
			} else if cursor == first && op == 'U' { // 一般向上翻页
				cursor--
				first--
			} else if cursor == (first+3) && op == 'D' { // 一般向下翻页
				cursor++
				first++
			} else if op == 'U' { // 其他情况，不翻页，只移动光标
				cursor--
			} else {
				cursor++
			}
		}
		for i := first; i <= first+3; i++ {
			fmt.Printf("%d ", i)
		}
		fmt.Println()
		fmt.Println(cursor)
	}
}

// 滑动窗口
// 显示的页面就是一个窗口，同时对向上和向下操作进行优化。首先我们知道如果在0、1、2、3这四个数中循环，
// 假设当前位置为i，光标向前移动一格后的位置为(i-1+n)%n；光标向后移动一格后的位置为(i+1)%n。
// 这个理论可以推广到1、2、3、4这四个数中循环，假设当前位置为i，光标向前移动一格后的位置为(i-1-1+n)%n+1；
// 光标向后移动一格后的位置为i%n+1。因此我们可以简化判断，直接用取余计算。

// 如果移动后的光标不在窗口内，则需要滑动窗口。如果光标在窗口起始位置前，则把窗口的起始位置和光标对齐；
// 如果光标在窗口后，则把窗口的末位置和光标对齐。
func method2() {
	var n int
	var ops string
	fmt.Scan(&n, &ops)
	cursor := 1         // 当前光标所在位置
	winBegin := 1       // 窗口的起始位置
	winEnd := min(4, n) // 窗口的结束位置
	for _, op := range ops {
		if op == 'U' {
			cursor = (cursor-1-1+n)%n + 1
		} else {
			cursor = cursor%n + 1
		}
		if cursor < winBegin { // 如果当前光标在窗口前，则将窗口往前移动
			winBegin = cursor
			winEnd = winBegin + 3
		} else if cursor > winEnd { // 如果当前光标在窗口后，则将窗口往后移动
			winEnd = cursor
			winBegin = winEnd - 3
		}
	}
	for i := winBegin; i <= winEnd; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Println(cursor)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
