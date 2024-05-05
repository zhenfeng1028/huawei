package main

import "fmt"

func main() {
	var m, n int
	var candidates, votes []string
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		var tmp string
		fmt.Scan(&tmp)
		candidates = append(candidates, tmp)

	}
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Scan(&tmp)
		votes = append(votes, tmp)
	}
	hashmap := make(map[string]int)
	for _, cand := range candidates {
		hashmap[cand] = 0
	}
	invalid := 0
	for _, vote := range votes {
		if _, ok := hashmap[vote]; ok {
			hashmap[vote]++
		} else {
			invalid++
		}
	}
	for _, v := range candidates {
		fmt.Printf("%s : %d\n", v, hashmap[v])
	}
	fmt.Printf("Invalid : %d\n", invalid)
}
