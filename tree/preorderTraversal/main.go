package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

type Command struct {
	Str  string
	root *BNode
}

func perderTraversal(root *BNode) (vals []int) {
	if root == nil {
		return nil
	}
	var stack []Command
	stack = append(stack, Command{
		"go",
		root,
	})
	for len(stack) > 0 {
		cmd := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cmd.Str == "print" {
			vals = append(vals, cmd.root.Data.(int))
		} else {
			if cmd.root.RightChild != nil {
				stack = append(stack, Command{
					"go",
					cmd.root.RightChild,
				})
			}
			if cmd.root.RightChild != nil {
				stack = append(stack, Command{
					"go",
					cmd.root.RightChild,
				})
			}
			stack = append(stack, Command{
				"print",
				cmd.root,
			})

		}

	}
	return vals
}

type Level struct {
	root *BNode
	L    int
}

func levelOrder(root *BNode) [][]int {
	var queue []Level
	res := make([][]int, 0)
	queue = append(queue, Level{
		root: root,
		L:    0,
	})
	for len(queue) > 0 {
		cmd := queue[0]
		nod := cmd.root
		level := cmd.L
		queue = queue[1:]
		if len(res) == cmd.L {
			res = append(res, []int{})
		}
		res[level] = append(res[level], nod.Data.(int))
		if nod.LeftChild != nil {
			queue = append(queue, Level{
				nod.LeftChild,
				level + 1,
			})
		}
		if nod.RightChild != nil {
			queue = append(queue, Level{
				nod.RightChild,
				level + 1,
			})
		}
	}
	return res
}

func main() {
	data := []int{3, 9, 20, 7, 8, 15, 7}
	root := ArrayToTree(data, 0, len(data)-1)
	fmt.Println(levelOrder(root))
}
