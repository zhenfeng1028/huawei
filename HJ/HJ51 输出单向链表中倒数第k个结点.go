package main

import "fmt"

type Node struct {
	val  int
	Next *Node
}

func main() {
	for {
		var n int
		fmt.Scan(&n)
		nodes := []int{}
		for i := 0; i < n; i++ {
			var node int
			fmt.Scan(&node)
			nodes = append(nodes, node)
		}
		var k int
		fmt.Scan(&k)
		head := &Node{nodes[0], nil}
		prev := head
		for i := 1; i < len(nodes); i++ {
			prev.Next = &Node{nodes[i], nil}
			prev = prev.Next
		}
		count := 0
		prev = head
		for prev != nil {
			count++
			prev = prev.Next
		}
		i := 1
		for head != nil {
			if i == count-k+1 {
				fmt.Println(head.val)
			}
			i++
			head = head.Next
		}
	}
}
