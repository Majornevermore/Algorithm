package main

import (
	"fmt"
)

// 变种问题 完全背包问题：每个物品可以无限使用
// 变种问题 多重背包问题：每个物品不只有一个，有num【i】个
// 多维费用背包问题：要考虑物品的体积和重量两个维度

var memo [][]int

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func bestValue(w []int, v []int, index, c int) int {
	if index < 0 || c <= 0 {
		return 0
	}
	if memo[index][c] != 0 {
		return memo[index][c]
	}
	res := bestValue(w, v, index-1, c)
	if c >= w[index] {
		res = max(res, v[index]+bestValue(w, v, index-1, c-w[index]))
	}
	memo[index][c] = res
	return res
}

func knaspsack(w []int, v []int, c int) int {
	n := len(w)
	memo = make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, c+1)
	}
	return bestValue(w, v, n-1, c)
}

func knaspsack1(w []int, v []int, c int) int {
	n := len(w)
	memo = make([][]int, 2)
	for i := range memo {
		memo[i] = make([]int, c+1)
	}
	for i := 0; i <= c; i++ {
		var temp int
		if i >= w[0] {
			temp = w[0]
		} else {
			temp = 0
		}
		memo[0][i] = temp
	}
	for i := 2; i < n; i++ {
		for j := 0; j <= c; j++ {
			memo[i%2][j] = memo[(i-1)%2][j]
			if j >= w[i] {
				memo[i%2][j] = max(memo[i%2][j], memo[(i-1)%2][j]+v[i])
			}
		}
	}
	return memo[(n-1)%2][c]
}

func main() {
	var w = []int{1, 8, 8, 10, 2}
	var v = []int{1, 3, 3, 10, 3}
	fmt.Println(knaspsack1(w, v, 20))
}
