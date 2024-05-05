package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin) // 和NewReader不同在于scanner读取多行
	input.Scan()                        // 读取第一行
	str1 := input.Text()                // 获取这一行的数据
	input.Scan()                        // 读取第二行，题中只读取两行
	str2 := input.Text()
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)
	hashmap := make(map[byte]int)
	for i := 0; i < len(str1); i++ {
		hashmap[str1[i]]++
	}
	fmt.Println(hashmap[str2[0]])
}
