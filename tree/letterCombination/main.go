package main

import (
	"fmt"
)

var letters = [10]string{
	" ", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
}

var reslice []string

func findSubLetters(digits string, index int, res string) {
	if len(digits) == index {
		reslice = append(reslice, res)
		return
	}
	c := digits[index]
	letter := letters[c-'0']
	for i := range letter {
		findSubLetters(digits, index+1, res+string(letter[i]))
	}
}

func letterCombinations(digits string) []string {
	reslice = nil
	if digits == "" {
		return reslice
	}
	findSubLetters(digits, 0, "")
	return reslice
}

var judge []int

var r [][]int

func permuteVector(nums []int, index int, res []int) {
	if index == len(nums) {
		r = append(r, res)
		return
	}
	lens := len(nums)
	for i := 0; i < lens; i++ {
		if judge[i] != 1 {
			judge[i] = 1
			permuteVector(nums, index+1, res)
			var s int = len(res) - 1
			res = res[:s-1]
			judge[i] = 0
		}
	}
	return

}

func permute(nums []int) [][]int {
	r = nil
	judge = make([]int, len(nums))
	if len(nums) == 0 {
		return r
	}
	var p []int
	permuteVector(nums, 0, p)
	return r
}

var tagert = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var judgeboard [][]int
var maxX int
var maxY int

func inArea(x, y int) bool {
	return x <= maxX && y <= maxY && x >= 0 && y >= 0
}

func searchWord(board [][]byte, word string, index int, startx int, starty int) bool {
	if index == len(word)-1 {
		return board[startx][starty] == word[index]
	}
	if board[startx][starty] == word[index] {
		judgeboard[startx][startx] = 1
		for i := 0; i < 4; i++ {
			newX := startx + tagert[i][0]
			newY := startx + tagert[i][1]
			if inArea(newX, newY) && judgeboard[newX][newY] != 1 {

				if searchWord(board, word, index+1, newX, newY) {
					return true
				}
			}
		}
		judgeboard[startx][startx] = 0
	}
	return false
}

func exist(board [][]byte, word string) bool {
	maxX = len(board) - 1
	maxY = len(board[0]) - 1
	judgeboard = make([][]int, maxX)
	for i := range judgeboard {
		judgeboard[i] = make([]int, maxY)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if searchWord(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func main() {
	letterCombinations("231")
	fmt.Println()

}
