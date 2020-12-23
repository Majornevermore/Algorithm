package main

import . "github.com/isdamir/gotype"

//请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点。传入函数的唯一参数为 要被删除的节点 。
//
//
//
//现有一个链表 -- head = [4,5,1,9]，它可以表示为:
//
//
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/delete-node-in-a-linked-list

func deleteDuplicates(head *LNode) *LNode {
	if head == nil {
		return head
	}
	hash := map[int]bool{}
	var temp []int
	list := head
	for list != nil {
		n := len(temp)
		if len(temp) == 0 {
			temp = append(temp, list.Data.(int))
		} else {
			if list.Data.(int) == temp[n-1] {
				temp = temp[:n-1]

			} else if list.Data.(int) > temp[n-1] {
				temp = append(temp, list.Data.(int))
			}
		}
		list = list.Next
	}
	head = &LNode{}
	for i := 0; i < len(temp); i++ {
		head.Next = &LNode{Data: temp[i]}
	}
	return head.Next
}
func main() {
	head := &LNode{
		Data: 1,
	}
	head.Next = &LNode{
		Data: 1,
	}
	head.Next = &LNode{
		Data: 1,
	}
	deleteDuplicates(head)
}
