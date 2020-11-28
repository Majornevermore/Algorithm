package main

import (
	"fmt"

	. "github.com/isdamir/gotype"
)

func MoveBottomToTops(s *SliceStack) {
	if s.IsEmpty() {
		return
	}
	top1 := s.Pop()
	if !s.IsEmpty() {
		MoveBottomToTops(s)
		top2 := s.Pop()
		s.Push(top1)
		s.Push(top2)
	} else {
		s.Push(top1)
	}
}

func CreateStack(list []int) *SliceStack {
	stack := NewSliceStack()
	for _, v := range list {
		stack.Push(v)
	}
	return stack
}
func PrintStack(s *SliceStack) {
	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}
	fmt.Println()
}

func Reverse(s *SliceStack) {
	if s.IsEmpty() {
		return
	}
	MoveBottomToTops(s)
	tops := s.Pop()
	Reverse(s)
	s.Push(tops)
}
func main() {
	l := []int{1, 0, 3, 4, 5}

	stack := CreateStack(l)
	//PrintStack(stack)
	Reverse(stack)
	PrintStack(stack)
}
