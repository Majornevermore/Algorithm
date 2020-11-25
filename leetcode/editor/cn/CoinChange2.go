package main

import "strconv"

//给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。
//
//
//
//
//
//
// 示例 1:
//
// 输入: amount = 5, coins = [1, 2, 5]
//输出: 4
//解释: 有四种方式可以凑成总金额:
//5=5
//5=2+2+1
//5=2+1+1+1
//5=1+1+1+1+1
//
//
// 示例 2:
//
// 输入: amount = 3, coins = [2]
//输出: 0
//解释: 只用面额2的硬币不能凑成总金额3。
//
//
// 示例 3:
//
// 输入: amount = 10, coins = [10]
//输出: 1
//
//
//
//
// 注意:
//
// 你可以假设：
//
//
// 0 <= amount (总金额) <= 5000
// 1 <= coin (硬币面额) <= 5000
// 硬币种类不超过 500 种
// 结果符合 32 位符号整数
//
// 👍 272 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
var hashmap map[string]int

func dp2(coins []int, index, n int) int {
	key := strconv.Itoa(index) + "," + strconv.Itoa(n)

	ret := 0
	if v, ok := hashmap[key]; ok {
		return v
	}
	if index >= len(coins) {
		if n == 0 {
			return 1
		}
		return 0
	}
	for i := 0; i*coins[index] <= n; i++ {
		ret += dp2(coins, index+1, n-i*coins[index])
	}
	hashmap[key] = ret
	return ret
}

func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			temp := coins[i-1]
			if j-temp >= 0 {
				dp[j] = dp[j] + dp[j-temp]
			}
		}
	}
	return dp[amount]
}

// 递归解法时间超时
func change1(amount int, coins []int) int {
	hashmap = make(map[string]int)
	return dp2(coins, 0, amount)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
