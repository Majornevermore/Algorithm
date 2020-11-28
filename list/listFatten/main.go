package main

type LNodeStruct struct {
	data  int
	right *LNodeStruct
	down  *LNodeStruct
}

func (p *LNodeStruct) Instert(pHead *LNodeStruct, data int) *LNodeStruct {
	newList := &LNodeStruct{data: data, down: pHead}
	pHead = newList
	return pHead
}
func Merge(a, b *LNodeStruct) *LNodeStruct {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	var result *LNodeStruct
	if a.data < b.data {
		result = a
		result.down = Merge(a.down, b)
	} else {
		result = b
		result.down = Merge(a, b.down)
	}
	return result

}

func Fattern(root *LNodeStruct) *LNodeStruct {
	if root == nil || root.right == nil {
		return root
	}
	root.right = Fattern(root.right)
	// ba root 节点对应的链表和右边合并
	root = Merge(root, root.right)
	return root
}
func main() {

}
