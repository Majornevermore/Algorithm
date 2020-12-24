package main

import "fmt"

func LIS(arr []int) []int {
	maxLen := 0
	n := len(arr)
	var res []int
	if n == 0 {
		return res
	}
	res = make([]int, n)
	pos := make([]int, n)
	res[maxLen] = arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > res[maxLen] {
			maxLen++
			res[maxLen] = arr[i]
			pos[i] = maxLen
		} else {
			l, r := 0, maxLen
			for l <= r {
				mid := l + (r-l)/2
				if res[mid] < arr[i] {
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			if l != -1 {
				res[l] = arr[i]
				pos[i] = l
			} else {
				res[0] = arr[i]
				pos[i] = 0
			}
		}
	}
	ans := make([]int, maxLen+1)
	for i := n - 1; i >= 0; i-- {
		if pos[i] == maxLen {
			ans[maxLen] = arr[i]
			maxLen--
		}
	}
	return ans
}

func getMinStack(op [][]int) []int {

	var store []int
	var res []int
	var minSlice []int
	for i := 0; i < len(op); i++ {
		if op[i][0] == 1 {
			store = append(store, op[i][1])
			if len(minSlice) == 0 {
				minSlice = append(minSlice, op[i][1])
			} else {
				n := len(minSlice)
				temp := min(minSlice[n-1], op[i][1])
				minSlice = append(minSlice, temp)
			}
		} else if op[i][0] == 2 {
			n := len(store)
			if n > 0 {
				store = store[:n-1]
				minSlice = minSlice[:n-1]
			}
		} else if op[i][0] == 3 {
			n := len(minSlice)
			if n > 0 {
				res = append(res, minSlice[n-1])
			}
		}
	}
	return res
}

func trans(s string, n int) string {
	if n == 0 {
		return s
	}
	i, j := 0, n-1
	b := []byte(s)
	for i < j {
		temp := b[i]
		b[i] = b[j]
		b[j] = temp
		i++
		j--
	}
	for i := 0; i < n; i++ {
		j := i
		for j < n && b[j] != ' ' {
			j++
		}
		l, r := i, j-1
		for l < r {
			temp := b[l]
			b[l] = b[r]
			b[r] = temp
			l++
			r--
		}
		i = j
	}
	var res string
	for i := 0; i < n; i++ {
		if b[i] != ' ' && b[i] <= 'Z' && b[i] >= 'A' {
			temp := b[i] - 'A' + 'a'
			res += string(temp)
		} else if b[i] != ' ' && b[i] <= 'z' && b[i] >= 'a' {
			temp := b[i] - 'a' + 'A'
			res += string(temp)
		} else {
			res += " "
		}
	}
	return res
}

func solve(M int, N int) string {
	var ans string
	c := "0123456789ABCDEF"
	b := []byte(c)
	flag := 0
	if M < 0 {
		M = -M
		flag = 1
	}
	for M != 0 {
		r := b[M%N : M%N+1]
		ans = string(r) + ans
		M /= N
	}
	if flag == 1 {
		ans = "-" + ans
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func test() {
	var res [][]int
	var df func(a []int, n int)
	df = func(a []int, n int) {
		if n == 10 {
			res = append(res, a)
		}
		for i := 0; i < 10; i++ {
			a = append(a, i)
			df(a, i+1)
		}
	}
	df([]int{}, 0)
}

/**
 * longest common subsequence
 * @param s1 string字符串 the string
 * @param s2 string字符串 the string
 * @return string字符串
 */
func LCS(s1 string, s2 string) string {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = maxlc(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	if dp[m][n] == 0 {
		return "-1"
	}
	var res string
	i := m - 1
	j := n - 1
	for i >= 0 && j >= 0 {
		if s1[i] == s2[j] {
			res += string(s1[i])
			i--
			j--
		} else {
			if dp[i+1][j] >= dp[i][j+1] {
				j--
			} else {
				i--
			}
		}
	}
	return reverse(res)
}

func reverse(s string) string {
	b := []byte(s)
	left, right := 0, len(s)-1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
	return string(b)
}

func maxlc(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	b := LCS("1A2C6D4B560", "B1D23CA45B6A0")
	fmt.Println(b)
}
