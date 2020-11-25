package main

import (
	"fmt"
	"math"
)

//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回
// -1。
//
// 你可以认为每种硬币的数量是无限的。
//
//
//
// 示例 1：
//
//
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//
// 示例 2：
//
//
//输入：coins = [2], amount = 3
//输出：-1
//
// 示例 3：
//
//
//输入：coins = [1], amount = 0
//输出：0
//
//
// 示例 4：
//
//
//输入：coins = [1], amount = 1
//输出：1
//
//
// 示例 5：
//
//
//输入：coins = [1], amount = 2
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 231 - 1
// 0 <= amount <= 104
//
// Related Topics 动态规划
// 👍 932 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 递归法

func minNums(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func dp(coins []int, n int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}
	if memCoin[n] != -99 {
		return memCoin[n]
	}
	res := math.MaxInt32
	for i := range coins {
		subproblem := dp(coins, n-coins[i])
		if subproblem == -1 {
			continue
		}
		res = minNums(1+subproblem, res)
	}
	if res != math.MaxInt32 {
		memCoin[n] = res
	} else {
		memCoin[n] = -1
	}

	return memCoin[n]
}

var memCoin []int

//递归方法加记忆
func coinChangeDp(coins []int, amount int) int {
	memCoin = make([]int, amount+1)
	for i := range memCoin {
		memCoin[i] = -99
	}
	return dp(coins, amount)
}

// 动态规划，迭代
func coinChange(coins []int, amount int) int {
	mem := make([]int, amount+1)
	for i := range mem {
		mem[i] = amount + 1
	}
	mem[0] = 0

	for i := 1; i <= amount; i++ {
		for _, v := range coins {
			if i-v < 0 {
				continue
			}
			mem[i] = minNums(mem[i], mem[i-v]+1)
		}
	}
	if mem[amount] == amount+1 {
		return -1
	}
	return mem[amount]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

	a := []int{1, 3, 65, 77}
	for i, v := range a {
		fmt.Println(i, v)
	}

}
