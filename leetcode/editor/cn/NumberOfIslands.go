package main

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//
//
// 示例 1：
//
//
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
//
//
// 示例 2：
//
//
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'
//
// Related Topics 深度优先搜索 广度优先搜索 并查集
// 👍 860 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
var d = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
var m, n int
var visited [][]int

func InArea(x, y int) bool {
	return x < m && y < n && x >= 0 && y >= 0
}

func dfs(grid [][]byte, startX, startY int) {
	visited[startX][startY] = 1
	for i := 0; i < 4; i++ {
		newX := startX + d[i][0]
		newY := startY + d[i][1]
		if InArea(newX, newY) && visited[newX][newY] != 1 && grid[newX][newY] == '1' {
			dfs(grid, newX, newY)
		}
	}
	return
}
func numIslands(grid [][]byte) int {
	m = len(grid)
	if m == 0 {
		return 0
	}
	n = len(grid[0])
	visited = make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && visited[i][j] != 1 {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
