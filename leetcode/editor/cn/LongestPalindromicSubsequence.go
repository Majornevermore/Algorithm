package main

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
// 👍 341 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
var dpLongest [][]int

func longestPalindromeSubseq(s string) int {
	sl := len(s)
	dpLongest = make([][]int, sl)
	for i := 0; i < sl; i++ {
		dpLongest[i] = make([]int, sl)
	}
	// i > j 等于0 i== j时就是一个串所以肯定返回1
	for i := 0; i < sl; i++ {
		dpLongest[i][i] = 1
	}
	for i := sl - 1; i >= 0; i-- {
		for j := i + 1; j < sl; j++ {
			if s[i] == s[j] {
				dpLongest[i][j] = dpLongest[i+1][j-1] + 2
			} else {
				dpLongest[i][j] = maxLongest(dpLongest[i+1][j], dpLongest[i][j-1])
			}
		}
	}
	return dpLongest[0][sl-1]

}

func maxLongest(x, y int) int {
	if x < y {
		return y
	}

	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
