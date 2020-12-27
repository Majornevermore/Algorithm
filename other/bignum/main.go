package main

import (
	"fmt"
	"sort"
	"strconv"
)

func maxLength(arr []int) int {
	hasmap := make(map[int]bool)
	before, after := 0, 0
	n := len(arr)
	var maxlen int
	for before <= after && after < n {
		if _, ok := hasmap[arr[after]]; !ok {
			hasmap[arr[after]] = true
			after++
		} else {
			delete(hasmap, arr[before])
			before++
		}
		if len(hasmap) > maxlen {
			maxlen = len(hasmap)
		}
	}
	return maxlen
}

func jumpFloor(number int) int {
	var dp = make([]int, number+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= number; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[number]
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	hashmap := map[byte]byte{}
	hashmap['('] = ')'
	hashmap['{'] = '}'
	hashmap['['] = ']'
	i, j := 0, len(s)-1
	for i < j {
		if s[j] == hashmap[s[i]] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func solve(s string, t string) string {
	// write code here
	if len(s) > len(t) {
		temp := s
		s = t
		t = temp
	}
	short, long := len(s), len(t)
	var carry int
	var res string
	for i := 0; i < short; i++ {
		a := s[short-i-1] - '0'
		b := t[long-i-1] - '0'
		v := carry + int(a) + int(b)
		res = res + strconv.Itoa(v%10)
		carry = v / 10
	}
	for j := short; j < long; j++ {
		b := t[long-j-1] - '0'
		v := carry + int(b)
		res = res + strconv.Itoa(v%10)
		carry = v / 10
	}
	if carry != 0 {
		res = res + strconv.Itoa(carry)
	}
	return reverse(res)
}

func reverse(str string) string {
	var result string
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result = result + fmt.Sprintf("%c", str[strLen-i-1])
	}
	return result
}

func permuteUnique(num []int) [][]int {
	var res [][]int
	sort.Ints(num)
	n := len(num)
	if len(num) == 0 {
		return res
	}
	used := make([]bool, n)
	var search func([]int, []int)
	search = func(nums []int, ans []int) {
		if len(ans) == n {
			res = append(res, ans)
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue
			}
			used[i] = true
			ans = append(ans, nums[i])
			search(nums, ans)
			used[i] = false
			ans = ans[:len(ans)-1]
		}
	}
	var ans []int
	search(num, ans)
	return res
}

type Tree struct {
	val   int
	Left  *Tree
	Right *Tree
}

func solve1(xianxu []int, zhongxu []int) []int {
	// write code here
	p := buildTree(xianxu, zhongxu, 0, len(xianxu)-1, 0, len(zhongxu)-1)
	var res []int
	if p == nil {
		return res
	}
	var q []*Tree
	q = append(q, p)
	for len(q) > 0 {
		n := len(q)
		temp := q[0]
		res = append(res, temp.val)
		for i := 0; i < n; i++ {
			temp = q[0]
			q = q[:1]
			if temp.Right != nil {
				q = append(q, temp.Right)
			}
			if temp.Left != nil {
				q = append(q, temp.Right)
			}
		}
	}
	return res
}

func buildTree(xianxu, zhongxu []int, le1, ri1, le2, ri2 int) *Tree {
	if le1 > ri1 || le1-ri1 != le2-ri2 {
		return nil
	}
	index := le2
	for zhongxu[index] != xianxu[le1] {
		index++ // 在中序遍历中找到第一个根节点（前序遍历的第一个数）
	}
	p := &Tree{val: xianxu[le1]}
	p.Left = buildTree(xianxu, zhongxu, le1+1, le1+index-le2, le2, index-1)
	p.Right = buildTree(xianxu, zhongxu, le1+index-le2+1, ri1, index+1, ri2)
	return p
}

func main() {
	//jumpFloor(4)
	//fmt.Println(maxLength([]int{2, 2, 3, 4, 3}))
	//fmt.Println(solve("1292", "123"))
	//fmt.Print(isValid("([])"))
	permuteUnique([]int{1, 2, 3, 1})
}
