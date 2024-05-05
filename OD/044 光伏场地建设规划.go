package main

import "fmt"

var m, n, l, gc int // 长、宽、正方形边长、发电量
var arr [][]int

func main() {
	fmt.Scan(&m, &n, &l, &gc)
	for i := 0; i < m; i++ {
		var row []int
		for j := 0; j < n; j++ {
			var tmp int
			fmt.Scan(&tmp)
			row = append(row, tmp)
		}
		arr = append(arr, row)
	}
	cnt := 0
	for i := 0; i <= m-l; i++ {
		for j := 0; j <= n-l; j++ {
			sum := 0
			for p := i; p < i+l; p++ {
				for q := j; q < j+l; q++ {
					sum += arr[p][q]
				}
			}
			if sum >= gc {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
