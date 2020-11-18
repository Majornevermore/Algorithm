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
// 👍 861 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
var d = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var visited [][]int
var xMax, yMax int

func inArea(x, y int) bool {
	return x < xMax && y < yMax && x >= 0 && y >= 0
}

func dfs(grid [][]byte, startX, startY int) {
	visited[startX][startY] = 1
	for i := 0; i < 4; i++ {
		newX := startX + d[i][0]
		newY := startY + d[i][1]
		if inArea(newX, newY) && grid[newX][newY] == '1' && visited[newX][newY] != 1 {
			dfs(grid, newX, newY)
		}
	}
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	xMax = len(grid)
	yMax = len(grid[0])
	var res int
	visited = make([][]int, xMax)
	for i := range visited {
		visited[i] = make([]int, yMax)
	}
	for i := 0; i < xMax; i++ {
		for j := 0; j < yMax; j++ {
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
