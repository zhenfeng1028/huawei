package main

import "fmt"

func main() {
	var str string
	var n int
	fmt.Scan(&str)
	fmt.Scan(&n)
	hashmap := make(map[float64][]string)
	for i := 0; i <= len(str)-n; i++ {
		subStr := str[i : i+n]
		r := GCratio(subStr)
		if _, ok := hashmap[r]; !ok {
			hashmap[r] = make([]string, 0)
		}
		hashmap[r] = append(hashmap[r], subStr)
	}
	maxRatio := 0.0
	for k, _ := range hashmap {
		if k > maxRatio {
			maxRatio = k
		}
	}
	fmt.Println(hashmap[maxRatio][0])
}

func GCratio(s string) float64 {
	count := 0
	for _, c := range s {
		if c == 'C' || c == 'G' {
			count++
		}
	}
	return float64(count) / float64(len(s))
}
