package main
//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
//
// 
//
//说明：
//
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
//你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-sorted-array

func merge(nums1 []int, m int, nums2 []int, n int)  {
	n1right , n2right, total := m -1, n-1, m+n-1
	for n1right >= 0 && n2right >= 0 {
		if nums1[n1right] > nums2[n2right] {
			nums1[total] = nums1[n1right]
			n1right--
		} else {
			nums1[total] = nums2[n2right]
			n2right--
		}
		total--
	}
	for n2right >= 0 {
		nums1[total] = nums2[n2right]
		total--
		n2right--
	}
}

func main()  {
	
}
