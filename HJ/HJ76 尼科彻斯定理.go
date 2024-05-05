package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)
	if m == 1 {
		fmt.Println(1)
	} else {
		for i := 0; i < m-1; i++ {
			fmt.Printf("%d+", m*m-(m-1)+i*2)
		}
		fmt.Println(m*m - (m - 1) + (m-1)*2)
	}
}
