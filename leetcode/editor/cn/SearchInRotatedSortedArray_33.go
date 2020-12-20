package main
<<<<<<< HEAD

//升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为 [4,5,6,7,0,1,2] ）。
//
//
// 请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
//
//
// 示例 2：
//
//
//输入：nums = [4,5,6,7,0,1,2], target = 3
//输出：-1
//
// 示例 3：
//
//
//输入：nums = [1], target = 0
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5000
// -10^4 <= nums[i] <= 10^4
// nums 中的每个值都 独一无二
// nums 肯定会在某个点上旋转
// -10^4 <= target <= 10^4
//
// Related Topics 数组 二分查找
// 👍 1104 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 && nums[0] != target {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if (nums[mid] > target && target >= nums[left]) ||
			(nums[mid] <= nums[right] && (target > nums[right] || target < nums[mid])) {
=======
//给你一个整数数组 nums ，和一个整数 target 。 
//
// 该整数数组原本是按升序排列，但输入时在预先未知的某个点上进行了旋转。（例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2]
// ）。 
//
// 请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。 
// 
//
// 示例 1： 
//
// 
//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
// 
//
// 示例 2： 
//
// 
//输入：nums = [4,5,6,7,0,1,2], target = 3
//输出：-1 
//
// 示例 3： 
//
// 
//输入：nums = [1], target = 0
//输出：-1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 5000 
// -10^4 <= nums[i] <= 10^4 
// nums 中的每个值都 独一无二 
// nums 肯定会在某个点上旋转 
// -10^4 <= target <= 10^4 
// 
// Related Topics 数组 二分查找 
// 👍 1101 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 && nums[0] == target{
		return 0
	}
	left , right := 0, n -1
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// 判断是否在前半部分查找
		if (nums[left] <= target && target <= nums[mid]) ||
			(nums[mid] <= nums[right] && (target < nums[mid] || target > nums[right]))  {
>>>>>>> origin/master
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
<<<<<<< HEAD

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
=======
//leetcode submit region end(Prohibit modification and deletion)


func main() {
	search([]int {4,5,6,7,0,1,2}, 0)
}
>>>>>>> origin/master
