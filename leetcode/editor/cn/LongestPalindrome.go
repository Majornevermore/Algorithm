package main

import "fmt"

//给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
//
// 在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
//
// 注意:
//假设字符串的长度不会超过 1010。
//
// 示例 1:
//
//
//输入:
//"abccccdd"
//
//输出:
//7
//
//解释:
//我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
//
// Related Topics 哈希表
// 👍 248 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
var hasMap map[byte]int

func longestPalindrome(s string) int {
	l := len(s)
	var res int
	hasMap := make(map[byte]int, l)
	for i := 0; i < l; i++ {
		hasMap[s[i]]++
	}
	for _, v := range hasMap {
		res += v / 2 * 2
		if v%2 == 1 && res%2 == 0 {
			res++
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))
}
