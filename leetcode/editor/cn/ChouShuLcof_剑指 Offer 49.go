package main

//我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
//
//
//
// 示例:
//
// 输入: n = 10
//输出: 12
//解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
//
// 说明:
//
//
// 1 是丑数。
// n 不超过1690。
//
//
// 注意：本题与主站 264 题相同：https://leetcode-cn.com/problems/ugly-number-ii/
// Related Topics 数学
// 👍 96 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func nthUglyNumber(n int) int {
	// dp[i] 为 i+1个丑数
	//第一个丑数为1
	if n < 1 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
	a, b, c := 0, 0, 0
	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = minTree(n2, n3, n5)
		if n2 == dp[i] {
			a++
		}
		if n3 == dp[i] {
			b++
		}
		if n5 == dp[i] {
			c++
		}
	}
	return dp[n-1]
}

func minTree(a, b, c int) int {
	if a < minTwo(b, c) {
		return a
	}
	return minTwo(b, c)
}

func minTwo(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
