package main

//给定一个无序的整数数组，找到其中最长上升子序列的长度。
//
// 示例:
//
// 输入: [10,9,2,5,3,7,101,18]
//输出: 4
//解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
//
// 说明:
//
//
// 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
// 你算法的时间复杂度应该为 O(n2) 。
//
//
// 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
// Related Topics 二分查找 动态规划
// 👍 1165 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func max1(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func lengthOfLIS(nums []int) int {
	meme := make([]int, len(nums))
	for i := range meme {
		meme[i] = 1
	}
	res := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				meme[i] = max1(meme[i], meme[j]+1)
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		if res < meme[i] {
			res = meme[i]
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
