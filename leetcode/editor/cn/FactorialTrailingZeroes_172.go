package main

//给定一个整数 n，返回 n! 结果尾数中零的数量。
//
// 示例 1:
//
// 输入: 3
//输出: 0
//解释: 3! = 6, 尾数中没有零。
//
// 示例 2:
//
// 输入: 5
//输出: 1
//解释: 5! = 120, 尾数中有 1 个零.
//
// 说明: 你算法的时间复杂度应为 O(log n) 。
// Related Topics 数学
// 👍 389 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func trailingZeroes(n int) int {
	var res int
	delive := 5
	for delive <= n {
		res += n / delive
		delive *= 5
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	trailingZeroes(125)
}
