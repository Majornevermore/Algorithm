package main

//给定一个32位整数 num，你可以将一个数位从0变为1。请编写一个程序，找出你能够获得的最长的一串1的长度。
//
// 示例 1：
//
// 输入: num = 1775(110111011112)
//输出: 8
//
//
// 示例 2：
//
// 输入: num = 7(01112)
//输出: 4
//
// Related Topics 位运算
// 👍 20 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func reverseBits(num int) int {
	var count int
	for num&(num-1) != 0 {
		count++
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
