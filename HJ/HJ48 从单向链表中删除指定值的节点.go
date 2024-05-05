package main

import "fmt"

type Node struct {
	val  int
	Next *Node
}

// 6 2 1 2 3 2 5 1 4 5 7 2 2
func main() {
	var n int
	fmt.Scan(&n)
	nodes := []int{}
	for i := 0; i < (n-1)*2+2; i++ {
		var node int
		fmt.Scan(&node)
		nodes = append(nodes, node)
	}
	head := &Node{nodes[0], nil}
	deleteVal := nodes[len(nodes)-1]
	remaining := nodes[1 : len(nodes)-1]
	for i := 0; i < len(remaining); i = i + 2 {
		prev := head
		tmp := &Node{remaining[i], nil}
		for prev != nil {
			if prev.val == remaining[i+1] {
				tmp.Next = prev.Next
				prev.Next = tmp
				break
			}
			prev = prev.Next
		}
	}
	slow := &Node{0, head}
	fast := head
	preHead := slow
	for fast != nil {
		if fast.val == deleteVal {
			slow.Next = fast.Next
		}
		slow = slow.Next
		fast = fast.Next
	}
	head = preHead.Next
	for head != nil {
		fmt.Printf("%d ", head.val)
		head = head.Next
	}
}
