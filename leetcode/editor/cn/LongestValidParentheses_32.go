package main

//给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
//
// 示例 1:
//
// 输入: "(()"
//输出: 2
//解释: 最长有效括号子串为 "()"
//
//
// 示例 2:
//
// 输入: ")()())"
//输出: 4
//解释: 最长有效括号子串为 "()()"
//
// Related Topics 字符串 动态规划
// 👍 1119 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == ')' && i-1 >= 0 && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
			if i-dp[i-1]-2 >= 0 {
				dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
			} else {
				dp[i] = dp[i-1] + 2
			}

		}
	}
	var max int
	for i := 0; i < len(dp); i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
