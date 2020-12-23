package main

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
// 👍 219 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) (ans int) {
	columns := make([]bool, n)        // 列上是否有皇后
	diagonals1 := make([]bool, 2*n-1) // 左上到右下是否有皇后
	diagonals2 := make([]bool, 2*n-1) // 右上到左下是否有皇后
	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			ans++
			return
		}
		for col, hasQueen := range columns {
			d1, d2 := row+n-1-col, row+col
			if hasQueen || diagonals1[d1] || diagonals2[d2] {
				continue
			}
			columns[col] = true
			diagonals1[d1] = true
			diagonals2[d2] = true
			backtrack(row + 1)
			columns[col] = false
			diagonals1[d1] = false
			diagonals2[d2] = false
		}
	}
	backtrack(0)
	return
}

//
//if n == 0 {
//	return 0
//}
//var ans int
//cow := make([]int, n)
//dia1 := make([]int, 2*n-1)
//dia2 := make([]int, 2*n-1)
//var dfs func(int)
//dfs = func(r int) {
//	if r == n {
//		ans++
//		return
//	}
//	for c, hasQ := range cow {
//		if hasQ == 1 || dia1[c+r] == 1 || dia1[r-c+n-1] == 1 {
//			continue
//		}
//		cow[c] = 1
//		dia1[c+r] = 1
//		dia2[r-c+n-1] = 1
//		dfs(r + 1)
//		cow[c] = 0
//		dia1[c+r] = 0
//		dia2[r-c+n-1] = 0
//	}
//}
//dfs(0)
//return ans

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
