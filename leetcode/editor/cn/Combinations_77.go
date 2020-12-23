package main

//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
// 示例:
//
// 输入: n = 4, k = 2
//输出:
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]
// Related Topics 回溯算法
// 👍 459 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func combine(n int, k int) [][]int {
	var res [][]int
	if n == 0 || k == 0 || k > n {
		return res
	}
	var dfsCombine func(start int, record []int)
	dfsCombine = func(start int, record []int) {
		if len(record) == k {
			// 栈上分配的对象，这个对象一直没变过
			temp := make([]int, len(record))
			copy(temp, record)
			res = append(res, record)
			return
		}
		for i := start; i <= n; i++ {
			record = append(record, i)
			dfsCombine(i+1, record)
			record = record[:len(record)-1]
		}
		return
	}
	dfsCombine(1, []int{})
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	combine(4, 2)
}
