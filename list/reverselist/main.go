package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

func ReverseList(node* LNode)  {
	if node == nil || node.Next == nil {
		return
	}
	var pre *LNode
	var cur *LNode
	next := node.Next
	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	node.Next = pre
}

func RecursiveChild(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := RecursiveChild(node)
	node.Next.Next = node
	node.Next = nil
	return  newHead
}

func Recursive(node *LNode)  {
	firsetNode := node.Next
	newHead := RecursiveChild(firsetNode)
	node.Next = newHead
}
func InsertReverseList(node* LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var cur *LNode
	var next *LNode
	cur = node.Next.Next
	fmt.Println(cur, node.Next.Next)
	node.Next.Next = nil
	fmt.Println(cur, node.Next.Next)
	for cur != nil {
		next = cur.Next
		cur.Next = node.Next
		node.Next = cur
		cur = next
	}
}

func Modify(p *int) {
	fmt.Printf("p value %d address %p\n", *p, &p)
	*p = 1
}


func ModiyMap(p map[string]int) {
	fmt.Printf("函数接受的map内存地址是%p\n", &p)
	p["zhangsan"] = 10
}

func ModiyStruct(p Person) {
	fmt.Printf("函数接受的struct内存地址是%p\n", &p)
	p.Name = "zhagnsna"
}

func ModiyStructPoint(p *Person) {
	fmt.Printf("函数接受的struct内存地址是%p\n", &p)
	p.Name = "zhagnsna"
}

type Person struct {
	Name string
}

func main() {
	s := &LNode{}
	CreateNode(s, 8)

	PrintNode("翻转前", s)
	InsertReverseList(s)
	PrintNode("翻转后的", s)
	//v := 2
	//p := &v
	//fmt.Printf("jin ru %d address %p\n", *p, &p)
	//Modify(p)
	//fmt.Printf("modify after %d address %p\n", *p, &p)
	//s := make (map[string]int)
	//s["zhangsan"] = 1
	//fmt.Printf("jinru zhiqian %p, %d\n", &s, s["zhangsan"])
	//ModiyMap(s)
	//fmt.Printf("jinru zhihou %p, %d\n", &s, s["zhangsan"])
	//
	//p := Person{"lisi"}
	//fmt.Printf("jinru struct %p, %s\n", &p, p.Name)
	//ModiyStruct(p)
	//fmt.Printf("jinru hou struct %p, %s\n", &p, p.Name)
	//
	//ptr := &Person{"lisis"}
	//fmt.Printf("jinru struct %p, %s\n", &ptr, ptr.Name)
	//ModiyStructPoint(ptr)
	//fmt.Printf("jinru hou struct %p, %s\n", &ptr, ptr.Name)
}
