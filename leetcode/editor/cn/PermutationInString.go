package main

//给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
//
// 换句话说，第一个字符串的排列之一是第二个字符串的子串。
//
// 示例1:
//
//
//输入: s1 = "ab" s2 = "eidbaooo"
//输出: True
//解释: s2 包含 s1 的排列之一 ("ba").
//
//
//
//
// 示例2:
//
//
//输入: s1= "ab" s2 = "eidboaoo"
//输出: False
//
//
//
//
// 注意：
//
//
// 输入的字符串只包含小写字母
// 两个字符串的长度都在 [1, 10,000] 之间
//
// Related Topics 双指针 Sliding Window
// 👍 202 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	var left int
	var right int
	need := make(map[byte]int)
	window := make(map[byte]int)
	var visable int
	for i := range s1 {
		need[s1[i]]++
	}
	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				visable++
			}
		}
		for right-left >= len(s1) {
			c1 := s2[left]
			left++
			if visable == len(need) {
				return true
			}
			if _, ok := need[c1]; ok {
				if window[c1] == need[c1] {
					visable--
				}
				window[c1]--
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
