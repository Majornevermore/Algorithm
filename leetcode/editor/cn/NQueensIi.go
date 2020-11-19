package main

import "strings"

//n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
//
//
// 上图为 8 皇后问题的一种解法。
//
// 给定一个整数 n，返回 n 皇后不同的解决方案的数量。
//
// 示例:
//
// 输入: 4
//输出: 2
//解释: 4 皇后问题存在如下两个不同的解法。
//[
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // 解法 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]
//
//
//
//
// 提示：
//
//
// 皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一或 N
//-1 步，可进可退。（引用自 百度百科 - 皇后 ）
//
// Related Topics 回溯算法
// 👍 211 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
var dia1 []int
var dia2 []int
var col []int

func dfss(n int, index int, bd [][]string, res *[][]string) {
	if index == n {
		temp := make([]string, n)
		for i := range temp {
			temp[i] = strings.Join(bd[i], "")
		}
		*res = append(*res, temp)
		return
	}
	for i := 0; i < n; i++ {
		if col[i] != 1 && dia1[index+i] != 1 && dia2[index-i+n-1] != 1 {
			bd[i][index] = "Q"
			col[i] = 1
			dia1[index+i] = 1
			dia2[index-i+n-1] = 1
			dfss(n, index+1, bd, res)
			bd[i][index] = "."
			col[i] = 0
			dia1[index+i] = 0
			dia2[index-i+n-1] = 0
		}
	}
}

func SolveNQueens(n int) [][]string {
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}
	dia1 = make([]int, 2*n-1)
	dia2 = make([]int, 2*n-1)
	col = make([]int, n)
	var result [][]string
	dfss(n, 0, bd, &result)
	return result
}

func totalNQueens(n int) int {
	return len(SolveNQueens(n))
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
