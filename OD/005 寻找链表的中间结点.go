package main

import "fmt"

type Node struct {
	Address int
	Data    int
	Next    int
}

func main() {
	var hAddr, N int
	fmt.Scan(&hAddr, &N)
	nodes := []Node{}
	for i := 0; i < N; i++ {
		node := Node{}
		fmt.Scan(&node.Address, &node.Data, &node.Next)
		nodes = append(nodes, node)
	}
	head := Node{}
	for _, node := range nodes {
		if node.Address == hAddr {
			head = node
			break
		}
	}
	datas := []int{}
	datas = append(datas, head.Data)
	next := head.Next
	for next != -1 {
		for _, node := range nodes {
			if node.Address == next {
				datas = append(datas, node.Data)
				next = node.Next
				break
			}
		}
	}
	fmt.Println(datas[len(datas)/2])
}
