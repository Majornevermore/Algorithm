package main

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。
//
// 返回滑动窗口中的最大值。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
//输入：nums = [1], k = 1
//输出：[1]
//
//
// 示例 3：
//
//
//输入：nums = [1,-1], k = 1
//输出：[1,-1]
//
//
// 示例 4：
//
//
//输入：nums = [9,11], k = 2
//输出：[11]
//
//
// 示例 5：
//
//
//输入：nums = [4,-2], k = 2
//输出：[4]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// 1 <= k <= nums.length
//
// Related Topics 堆 Sliding Window
// 👍 683 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) <= 0 {
		return nil
	}
	var (
		windows []int
		res     []int
	)
	//遍历数组
	for i, v := range nums {
		//i-k = 当窗口满时的最小下标
		//i-k >= 窗口最左边的数，则代表窗口已经满了
		if i >= k && windows[0] <= i-k {
			windows = windows[1:]
		}

		for {
			//如果窗口最后一个小于v，则去掉。为了维护了每次窗口最左边值最大
			if len(windows) > 0 && nums[windows[len(windows)-1]] <= v {
				windows = windows[:len(windows)-1]
			} else {
				break
			}
		}

		windows = append(windows, i)

		if i >= k-1 {
			res = append(res, nums[windows[0]])
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
