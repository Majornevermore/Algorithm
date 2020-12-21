package main

import "sort"

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复
//的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例：
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
//
// Related Topics 数组 双指针
// 👍 2829 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func twoSum2(nums []int, left, target int) [][]int {
	n := len(nums)
	var res [][]int
	right := n - 1
	for left < right {
		sum := nums[left] + nums[right]
		lv, rv := nums[left], nums[right]
		var ss []int
		if sum == target {
			ss = append(ss, lv, rv)
			res = append(res, ss)
			for left < right && nums[right] == rv {
				right--
			}
			for left < right && nums[left] == lv {
				left++
			}
		} else if sum > target {
			for left < right && nums[right] == rv {
				right--
			}
		} else if sum < target {
			for left < right && nums[left] == lv {
				left++
			}
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		re := twoSum2(nums, i+1, -num)

		for j := 0; j < len(re); j++ {
			re[j] = append(re[j], num)
			res = append(res, re[j])
		}

		for i < len(nums)-1 && num == nums[i+1] {
			i++
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	threeSum([]int{0, -1, 1})
}
