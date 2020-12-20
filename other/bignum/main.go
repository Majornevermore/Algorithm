package main

import (
	"fmt"
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

func main() {
	//jumpFloor(4)
	//fmt.Println(maxLength([]int{2, 2, 3, 4, 3}))
	//fmt.Println(solve("1292", "123"))
	fmt.Print(isValid("([])"))
}
