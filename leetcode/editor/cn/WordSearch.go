package main

import "fmt"

//给定一个二维网格和一个单词，找出该单词是否存在于网格中。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
// 示例:
//
// board =
//[
//  ['A','B','C','E'],
//  ['S','F','C','S'],
//  ['A','D','E','E']
//]
//
//给定 word = "ABCCED", 返回 true
//给定 word = "SEE", 返回 true
//给定 word = "ABCB", 返回 false
//
//
//
// 提示：
//
//
// board 和 word 中只包含大写和小写英文字母。
// 1 <= board.length <= 200
// 1 <= board[i].length <= 200
// 1 <= word.length <= 10^3
//
// Related Topics 数组 回溯算法
// 👍 683 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
var judgeboard [][]int
var maxX int
var maxY int
var tagert = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func inArea(x, y int) bool {
	return x < maxX && y < maxY && x >= 0 && y >= 0
}

func searchWord(board [][]byte, word string, index int, startx int, starty int) bool {
	if index == len(word)-1 {
		return board[startx][starty] == word[index]
	}
	if board[startx][starty] == word[index] {
		judgeboard[startx][starty] = 1
		for i := 0; i < 4; i++ {
			newX := startx + tagert[i][0]
			newY := starty + tagert[i][1]
			if inArea(newX, newY) && judgeboard[newX][newY] != 1 {

				if searchWord(board, word, index+1, newX, newY) {
					return true
				}
			}
		}
		judgeboard[startx][starty] = 0
	}
	return false
}

func exist(board [][]byte, word string) bool {
	maxX = len(board)
	maxY = len(board[0])
	judgeboard = make([][]int, maxX)
	for i := range judgeboard {
		judgeboard[i] = make([]int, maxY)
	}

	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			if searchWord(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCCED"))
}
