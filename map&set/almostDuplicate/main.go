package main

import (
	"math"
)

//220. 存在重复元素 III
//在整数数组 nums 中，是否存在两个下标 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值小于等于 t ，且满足 i 和 j 的差的绝对值也小于等于 ķ 。
//
//如果存在则返回 true，不存在返回 false。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/contains-duplicate-iii

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if math.Abs(float64(nums[j])-float64(nums[i])) <= float64(t) {
				return true
			}
		}
	}
	return false
}

func main() {

}
