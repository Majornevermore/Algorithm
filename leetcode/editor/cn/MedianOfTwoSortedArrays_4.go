package main

import "math"

//给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。
//
// 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
//
//
//
// 示例 1：
//
// 输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//
//
// 示例 2：
//
// 输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
//
// 示例 3：
//
// 输入：nums1 = [0,0], nums2 = [0,0]
//输出：0.00000
//
//
// 示例 4：
//
// 输入：nums1 = [], nums2 = [1]
//输出：1.00000
//
//
// 示例 5：
//
// 输入：nums1 = [2], nums2 = []
//输出：2.00000
//
//
//
//
// 提示：
//
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106
//
// Related Topics 数组 二分查找 分治算法
// 👍 3534 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		temp := nums1
		nums1 = nums2
		nums2 = temp
	}
	m = len(nums1)
	n = len(nums2)
	left, right := 0, m
	total := (m + n + 1) / 2
	//要满足 nums1[i-1] <= nums[j] && nums2[j-1] <= nums1[i]
	for left < right {
		i := left + (right-left+1)/2
		j := total - i
		if nums1[i-1] > nums2[j] {
			right = i - 1
		} else {
			left = i
		}
	}
	i := left
	j := total - i
	var nums1leftMax int
	var nums1rightMin int
	var nums2leftMax int
	var nums2rightmin int
	if i != 0 {
		nums1leftMax = nums1[i-1]
	} else {
		nums1leftMax = math.MinInt32
	}
	if i == m {
		nums1rightMin = math.MaxInt32
	} else {
		nums1rightMin = nums1[i]
	}
	if j == 0 {
		nums2leftMax = math.MinInt32
	} else {
		nums2leftMax = nums2[j-1]
	}
	if j == n {
		nums2rightmin = math.MaxInt32
	} else {
		nums2rightmin = nums2[j]
	}
	if ((m + n) % 2) == 1 {
		return float64(maxT(nums1leftMax, nums2leftMax))
	}
	temp := float64(maxT(nums1leftMax, nums2leftMax)) + float64(minT(nums2rightmin, nums1rightMin))
	temp = temp / 2
	return temp

}

func maxT(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func minT(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	findMedianSortedArrays([]int{1, 2}, []int{3, 4})
}
