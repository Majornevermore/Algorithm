package main

import "fmt"

//1、分析二分查找代码时，不要出现 else，全部展开成 else if 方便理解。
//
//2、注意「搜索区间」和 while 的终止条件，如果存在漏掉的元素，记得在最后检查。
//
//3、如需定义左闭右开的「搜索区间」搜索左右边界，只要在nums[mid] == target时做修改即可，搜索右侧时需要减一。
//
//4、如果将「搜索区间」全都统一成两端都闭，好记，只要稍改nums[mid] == target条件处的代码和返回的逻辑即可.

//第一个，最基本的二分查找算法：
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			return mid
		}
	}
	return -1
}

//第二个，寻找左侧边界的二分查找：
func leftSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			right = mid - 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left

}

//第三个，寻找右侧边界的二分查找：
func rightSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			left = mid + 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right

}

func main() {
	a := []int{1, 2, 2, 2, 4, 5}
	fmt.Println(leftSearch(a, 7))
	fmt.Println(rightSearch(a, 6))
	fmt.Println(binarySearch(a, 2))
}
