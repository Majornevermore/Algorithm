package main

import (
	. "github.com/isdamir/gotype"
)

func SwapPair(head *LNode) *LNode {
	p := head
	for p.Next != nil && p.Next.Next != nil {
		node1 := p.Next
		node2 := p.Next.Next
		next := node2.Next
		node2.Next = node1
		node1.Next = next
		p.Next = node2
		p = node1
	}
	return head
}

func main() {
	s := &LNode{}
	CreateNode(s, 8)
	s = SwapPair(s)
	PrintNode("交换后", s)
}
