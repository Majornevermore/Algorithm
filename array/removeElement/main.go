package main

import (
	"fmt"
)

// 26 27 给定个有序数组，对数组中的元素去重，使得原数组的每一个元素只有一个
func RemoveElement1(z []int, target int) []int{
	p ,s := -1, 0
	for i:=0; i<len(z); i++ {
		if z[i] == target {
			s = 1
			continue
		} else {
			p++
			if s == 1 {
				z[p] = z[i]
			}
			s = 0
		}

	}
	return z[:p+1]
}
// 80给定一个增序排列数组 nums ，你需要在 原地 删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。
//
//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii

func RemoveElement(z []int) []int{
	p, s := 1, 1
	for i:=0; i<len(z); i++ {
		if z[i] > z[p -1] {
			s = 1
		} else if z[i] == z[p-1] && s==1 {
			s = 0
		} else {
			continue
		}
		z[p] = z[i]
		p++
	}
	return z[:p]
}

func main()  {
	s:=[]int{3, 2, 2, 3}
	fmt.Println(RemoveElement1(s, 2))
}
