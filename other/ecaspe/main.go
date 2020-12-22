package main

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

func main() {
	solve(7, 2)
}
