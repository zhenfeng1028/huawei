package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	var length int
	fmt.Scan(&length)
	fmt.Println(str[:length])
}
