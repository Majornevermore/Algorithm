package main

import "strconv"

//给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
//
// 有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
// 和 "192.168@1.1" 是 无效的 IP 地址。
//
//
//
// 示例 1：
//
// 输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]
//
//
// 示例 2：
//
// 输入：s = "0000"
//输出：["0.0.0.0"]
//
//
// 示例 3：
//
// 输入：s = "1111"
//输出：["1.1.1.1"]
//
//
// 示例 4：
//
// 输入：s = "010010"
//输出：["0.10.0.10","0.100.1.0"]
//
//
// 示例 5：
//
// 输入：s = "101023"
//输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 3000
// s 仅由数字组成
//
// Related Topics 字符串 回溯算法
// 👍 473 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
const SEG_COUNT = 4

var (
	res      []string
	segSlice []int
)

func restoreIpAddresses(s string) []string {
	res = []string{}
	segSlice = make([]int, SEG_COUNT)

	dfs(s, 0, 0)
	return res
}

func dfs(s string, seg_index, start int) {

	if seg_index == SEG_COUNT {
		if start == len(s) {
			ipAddr := ""
			for i := 0; i < SEG_COUNT; i++ {
				ipAddr += strconv.Itoa(segSlice[i])
				if i != SEG_COUNT-1 {
					ipAddr += "."
				}
			}
			res = append(res, ipAddr)
		}
		return
	}
	if start == len(s) {
		return
	}
	if s[start] == '0' {
		segSlice[seg_index] = 0
		dfs(s, seg_index+1, start+1)
	}
	addr := 0
	for j := start; j < len(s); j++ {
		addr = addr*10 + int(s[j]-'0')
		if addr > 0 && addr <= 255 {
			segSlice[seg_index] = addr
			dfs(s, seg_index+1, j+1)
		} else {
			break
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
