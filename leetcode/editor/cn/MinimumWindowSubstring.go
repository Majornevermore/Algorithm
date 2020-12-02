package main

import (
	"fmt"
	"math"
)

//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
// 注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//
// 示例 1：
//
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//
//
// 示例 2：
//
//
//输入：s = "a", t = "a"
//输出："a"
//
//
//
//
// 提示：
//
//
// 1 <= s.length, t.length <= 105
// s 和 t 由英文字母组成
//
//
//
//进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？ Related Topics 哈希表 双指针 字符串 Sliding Window
// 👍 846 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	var left int
	var right int
	var vsiable int
	var start int
	var end int
	maxlen := math.MaxInt32
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	l := len(s)
	for right < l {
		if _, ok := need[s[right]]; ok {
			window[s[right]]++
			if window[s[right]] == need[s[right]] {
				vsiable++
			}
		}
		right++
		for vsiable == len(need) {
			if right-left < maxlen {
				start = left
				end = right
				maxlen = right - start
			}
			if _, ok := need[s[left]]; ok {
				if window[s[left]] == need[s[left]] {
					vsiable--
				}
				window[s[left]]--
			}
			left++
		}
	}
	return s[start:end]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	q := minWindow(s, t)
	fmt.Println(q)
}
