package main

import "fmt"

//给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000 。
//
//
//
// 示例 1:
//输入:
//
// "bbbab"
//
//
// 输出:
//
// 4
//
//
// 一个可能的最长回文子序列为 "bbbb"。
//
// 示例 2:
//输入:
//
// "cbbd"
//
//
// 输出:
//
// 2
//
//
// 一个可能的最长回文子序列为 "bb"。
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 只包含小写英文字母
//
// Related Topics 动态规划
// 👍 351 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max4(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func max4(i, j int) int {
	if i < j {
		return j
	}
	return i
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(longestPalindromeSubseq("abc1234321ab"))
}
