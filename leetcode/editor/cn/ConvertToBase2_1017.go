package main

//给出数字 N，返回由若干 "0" 和 "1"组成的字符串，该字符串为 N 的负二进制（base -2）表示。
//
// 除非字符串就是 "0"，否则返回的字符串中不能含有前导零。
//
//
//
// 示例 1：
//
// 输入：2
//输出："110"
//解释：(-2) ^ 2 + (-2) ^ 1 = 2
//
//
// 示例 2：
//
// 输入：3
//输出："111"
//解释：(-2) ^ 2 + (-2) ^ 1 + (-2) ^ 0 = 3
//
//
// 示例 3：
//
// 输入：4
//输出："100"
//解释：(-2) ^ 2 = 4
//
//
//
//
// 提示：
//
//
// 0 <= N <= 10^9
//
// Related Topics 数学
// 👍 33 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func baseNeg2(N int) string {
	return solve(N, -2)
}

func solve(M int, N int) string {
	var ans string
	c := "0123456789ABCDEF"
	b := []byte(c)
	flag := 0
	if M < 0 {
		M = -M
		flag = 1
	}
	for M != 0 {
		r := b[M%N : M%N+1]
		ans = string(r) + ans
		M /= N
	}
	if flag == 1 {
		ans = "-" + ans
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	baseNeg2(2)
}
