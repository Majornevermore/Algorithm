package main

import "github.com/isdamir/gotype"

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
func deleteNode(node *gotype.LNode) {
	if node == nil {
		return
	}
	if node.Next != nil {
		node.Data = node.Next.Data
		node.Next = node.Next.Next
	} else {
		node = nil
	}
}
func main()  {

}
