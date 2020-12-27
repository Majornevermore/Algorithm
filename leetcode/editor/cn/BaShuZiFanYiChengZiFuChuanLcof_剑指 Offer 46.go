package main

import "strconv"

//给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可
//能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
//
//
//
// 示例 1:
//
// 输入: 12258
//输出: 5
//解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
//
//
//
// 提示：
//
//
// 0 <= num < 231
//
// 👍 164 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func translateNum(num int) int {
	str := strconv.Itoa(num)
	dp := make([]int, len(str)+1)
	dp[0] = 1
	dp[1] = 1
	for i := 1; i < len(str); i++ {
		dp[i+1] = dp[i]
		cur := (str[i-1]-'0')*10 + str[i] - '0'
		if cur > 9 && cur <= 25 {
			dp[i+1] = dp[i-1] + dp[i]
		}
	}
	return dp[len(str)]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
