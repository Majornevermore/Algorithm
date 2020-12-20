package main

//给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
//
// 示例 1 :
//
//
//输入:nums = [1,1,1], k = 2
//输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
//
//
// 说明 :
//
//
// 数组的长度为 [1, 20,000]。
// 数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
//
// Related Topics 数组 哈希表
// 👍 720 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// pre[i] = {0....i}求和
// {j...i} 中存在k，则pre[i] - pre[j] = k 则pre[j] = pre[i] - k

func subarraySum(nums []int, k int) int {
	mp := map[int]int{}
	mp[0] = 1
	count, pre := 0, 0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := mp[pre-k]; ok {
			count += mp[pre-k]
		}
		// pre 前缀出现的次数
		mp[pre] += 1
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
