package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

//给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
//此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sort-colors

func sortColors(nums []int)  {
	s := make([]int, 3)
	for i := 0; i<len(nums); i++ {
		s[nums[i]]++
	}
	index := 0
	for i := 0; i<len(s); i++ {
		for j:=0; j<s[i]; j++ {
			nums[index] = i
			index++
		}
	}
}

func sortColorsNB(nums []int)  {
	zero := -1
	two := len(nums)
	for i := 0; i<two; {
		if nums[i] == 0 {
			i++
			zero++
		} else if nums[i] == 2 {
			gotype.SwapInt(nums, i, two-1)
			two--
		} else {
			zero++
			gotype.SwapInt(nums, zero, i)
		}
	}
}

func main()  {
	s := []int {2, 0, 2, 1, 1, 0}
	sortColors(s)
	fmt.Println(s)
}
