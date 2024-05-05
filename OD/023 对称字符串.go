package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	var arr [][2]int
	for i := 0; i < T; i++ {
		var n, k int
		fmt.Scan(&n, &k)
		arr = append(arr, [2]int{n, k})
	}
	res := []string{}
	for _, v := range arr {
		str := getSymmetricString(v[0])
		if str[v[1]] == 'B' {
			res = append(res, "blue")
		} else {
			res = append(res, "red")
		}
	}
	for _, v := range res {
		fmt.Println(v)
	}
}

func getSymmetricString(n int) string {
	if n == 1 {
		return "R"
	}
	last := getSymmetricString(n - 1)
	tmp := ""
	for i := 0; i < len(last); i++ {
		if last[i] == 'R' {
			tmp += "B"
		} else {
			tmp += "R"
		}
	}
	return tmp + last
}
