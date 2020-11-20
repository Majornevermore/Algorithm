package main

//给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
//
// 示例 1:
//
// 输入: 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1。
//
// 示例 2:
//
// 输入: 10
//输出: 36
//解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
//
// 说明: 你可以假设 n 不小于 2 且不大于 58。
// Related Topics 数学 动态规划
// 👍 407 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
var mem []int

func integerBreak1(n int) int {
	mem = make([]int, n+1)
	return breakInteger(n)
}
func max3(a, b, c int) int {
	temp := -1
	if a > b {
		temp = a
	} else {
		temp = b
	}
	if temp > c {
		return temp
	} else {
		return c
	}
}

//记忆法自顶向下来实现
func breakInteger(n int) int {
	if n == 1 {
		return 1
	}
	res := -1
	if mem[n] != 0 {
		return mem[n]
	}
	for i := 1; i <= n-1; i++ {
		res = max3(res, i*(n-i), i*breakInteger(n-i))
	}
	mem[n] = res
	return res
}

//动态规划，自底向上
func integerBreak(n int) int {
	mem = make([]int, n+1)
	mem[0] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i-1; j++ {
			mem[i] = max3(mem[i], j*(i-j), j*mem[i-j])
		}
	}
	return mem[n]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
