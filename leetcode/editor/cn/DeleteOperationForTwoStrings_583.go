package main

//给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。
//
//
//
// 示例：
//
// 输入: "sea", "eat"
//输出: 2
//解释: 第一步将"sea"变为"ea"，第二步将"eat"变为"ea"
//
//
//
//
// 提示：
//
//
// 给定单词的长度不超过500。
// 给定单词中的字符只含有小写字母。
//
// Related Topics 字符串
// 👍 164 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func finLongestStr(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dpMem := make([][]int, m+1)
	for i := range dpMem {
		dpMem[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dpMem[i][j] = 1 + dpMem[i-1][j-1]
			} else {
				dpMem[i][j] = maxD(dpMem[i-1][j], dpMem[i][j-1])
			}
		}
	}
	return dpMem[m][n]
}

func maxD(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	lcs := finLongestStr(word1, word2)
	return m - lcs + n - lcs
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
