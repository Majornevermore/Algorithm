package main

//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上
//被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//
//
//
// 示例 1：
//
// 输入：[1,2,3,1]
//输出：4
//解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 2：
//
// 输入：[2,7,9,3,1]
//输出：12
//解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 100
// 0 <= nums[i] <= 400
//
// Related Topics 动态规划
// 👍 1171 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
var mem []int

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func tryRob(nums []int, index int) int {
	if index >= len(nums) {
		return 0
	}
	res := 0
	if mem[index] != -1 {
		return mem[index]
	}
	for i := index; i < len(nums); i++ {
		res = max(res, nums[i]+tryRob(nums, i+2))
	}
	mem[index] = res
	return res
}

func rob(nums []int) int {
	mem = make([]int, len(nums))
	//return tryRob(nums, 0)
	n := len(nums)
	if n == 0 {
		return 0
	}
	mem[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		for j := i; j < n; j++ {
			var temp int
			if j+2 < n {
				temp = mem[j+2]
			} else {
				temp = 0
			}
			mem[i] = max(mem[i], temp+nums[j])
		}
	}
	return mem[0]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
