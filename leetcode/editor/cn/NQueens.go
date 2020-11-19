package main

import (
	"fmt"
	"strings"
)

//n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
//
//
// 上图为 8 皇后问题的一种解法。
//
// 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
//
// 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
// 示例：
//
// 输入：4
//输出：[
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // 解法 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]
//解释: 4 皇后问题存在两个不同的解法。
//
//
//
//
// 提示：
//
//
// 皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。
//
// Related Topics 回溯算法
// 👍 669 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

var dia1s []int
var dia2s []int
var cols []int

func dfsss(n int, index int, bd [][]string, res *[][]string) {
	if index == n {
		temp := make([]string, n)
		for i := range temp {
			temp[i] = strings.Join(bd[i], "")
		}
		*res = append(*res, temp)
		return
	}
	for i := 0; i < n; i++ {
		if cols[i] != 1 && dia1s[index+i] != 1 && dia2s[index-i+n-1] != 1 {
			bd[i][index] = "Q"
			cols[i] = 1
			dia1s[index+i] = 1
			dia2s[index-i+n-1] = 1
			dfsss(n, index+1, bd, res)
			bd[i][index] = "."
			cols[i] = 0
			dia1s[index+i] = 0
			dia2s[index-i+n-1] = 0
		}
	}
}

func solveNQueens(n int) [][]string {
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}
	dia1s = make([]int, 2*n-1)
	dia2s = make([]int, 2*n-1)
	cols = make([]int, n)
	var result [][]string
	dfsss(n, 0, bd, &result)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	b := solveNQueens(4)
	fmt.Println(b)
}
