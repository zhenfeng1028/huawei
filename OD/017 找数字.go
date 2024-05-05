package main

import (
	"fmt"
	"math"
	"sort"
)

var m, n int
var arr [][]int

func main() {
	fmt.Scan(&m, &n)
	for i := 0; i < m; i++ {
		var row []int
		for j := 0; j < n; j++ {
			var tmp int
			fmt.Scan(&tmp)
			row = append(row, tmp)
		}
		arr = append(arr, row)
	}
	hashmap := make(map[int][][2]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			val := arr[i][j]
			tmp := [2]int{i, j}
			if _, ok := hashmap[val]; !ok {
				hashmap[val] = make([][2]int, 0)
			}
			hashmap[val] = append(hashmap[val], tmp)
		}
	}
	res := [][]int{}
	for i := 0; i < m; i++ {
		tmp := []int{}
		for j := 0; j < n; j++ {
			val := arr[i][j]
			sameVal := hashmap[val]
			if len(sameVal) == 1 {
				tmp = append(tmp, -1)
				continue
			}
			distances := []int{}
			for _, v := range sameVal {
				distances = append(distances, distance([2]int{i, j}, v))
			}
			sort.Ints(distances)
			tmp = append(tmp, distances[1])
		}
		res = append(res, tmp)
	}
	fmt.Println(res)
}

func distance(p1, p2 [2]int) int {
	return int(math.Abs(float64(p1[0]-p2[0])) + math.Abs(float64(p1[1]-p2[1])))
}
