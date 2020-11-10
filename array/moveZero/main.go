package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

func swap(a, b int) (int, int)  {
	fmt.Println("shu ru", a, b)
	x, y := b, a
	fmt.Println("shu chu", x, y)
	return x, y
}

func MoveZero(z []int) [] int {
	k := 0
	for i:=0; i<len(z); i++ {
		if z[i] !=0 {
			if k != i {
				gotype.SwapInt(z, k, i)
				k++
			} else {
				k++
			}

		}
	}
	return  z
}

func main()  {
	s := []int {1, 2, 0, 4, 0, 6}
	MoveZero(s)
	fmt.Println(s)
}
