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

func main() {
	LIS([]int{2, 1, 5, 3, 6, 4, 8, 9, 7})
}
