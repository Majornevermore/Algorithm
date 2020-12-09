package main

//给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
// 注意:
//
//
// 每个数组中的元素不会超过 100
// 数组的大小不会超过 200
//
//
// 示例 1:
//
// 输入: [1, 5, 11, 5]
//
//输出: true
//
//解释: 数组可以分割成 [1, 5, 5] 和 [11].
//
//
//
//
// 示例 2:
//
// 输入: [1, 2, 3, 5]
//
//输出: false
//
//解释: 数组不能分割成两个元素和相等的子集.
//
//
//
// Related Topics 动态规划
// 👍 595 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
var memeCan [][]int

func tryPartition(nums []int, index, sum int) bool {
	if sum == 0 {
		return true
	}
	if sum < 0 || index < 0 {
		return false
	}
	if memeCan[index][sum] != -1 {
		return memeCan[index][sum] == 1
	}

	if tryPartition(nums, index-1, sum) ||
		tryPartition(nums, index-1, sum-nums[index]) {
		memeCan[index][sum] = 1
	} else {
		memeCan[index][sum] = 0
	}
	return memeCan[index][sum] == 1
}

func canPartitionDP(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	memeCan = make([][]int, len(nums))
	for i := range memeCan {
		memeCan[i] = make([]int, sum/2+1)
	}
	for i := 0; i < len(memeCan); i++ {
		for j := 0; j < len(memeCan[0]); j++ {
			memeCan[i][j] = -1
		}
	}
	return tryPartition(nums, len(nums)-1, sum/2)
}

func canPartition(nums []int) bool {
	var memdp [][]bool
	sum, max := 0, 0
	n := len(nums)
	if n < 2 {
		return false
	}

	for i := range nums {
		if max < nums[i] {
			max = nums[i]
		}
		sum += nums[i]
	}
	target := sum % 2
	if target != 0 {
		return false
	}
	target = sum / 2
	if max > target {
		return false
	}
	memdp = make([][]bool, n)
	for i := range memdp {
		memdp[i] = make([]bool, target+1)
	}
	for i := 0; i < n; i++ {
		memdp[i][0] = true
	}
	memdp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		v := nums[i]
		for j := 1; j <= target; j++ {
			if j < v {
				memdp[i][j] = memdp[i-1][j]
			} else {
				memdp[i][j] = memdp[i-1][j] || memdp[i-1][j-v]
			}
		}
	}
	return memdp[n-1][target]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
