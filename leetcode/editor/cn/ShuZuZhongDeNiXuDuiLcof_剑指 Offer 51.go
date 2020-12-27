package main

//在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
//
//
//
// 示例 1:
//
// 输入: [7,5,6,4]
//输出: 5
//
//
//
// 限制：
//
// 0 <= 数组长度 <= 50000
// 👍 287 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	temp := make([]int, len(nums))
	return _reversePairs(nums, temp)
}

const mod = 00001

func _reversePairs(nums []int, temp []int) int {
	ln := len(nums)
	if ln < 2 {
		return 0
	}
	mid := ln / 2
	left := nums[:mid]
	right := nums[mid:]
	n := _reversePairs(left, temp[:mid])
	n += _reversePairs(right, temp[mid:])
	l, r := 0, 0
	l1, lr := len(left), len(right)
	temp = temp[:0]
	for l < l1 && r < lr {
		if left[l] > right[r] {
			temp = append(temp, right[r])
			r++
			n += l1 - l
		} else {
			temp = append(temp, left[l])
			l++
		}
	}
	if l <= l1 {
		temp = append(temp, left[l:]...)
	} else {
		temp = append(temp, right[r:]...)
	}
	copy(nums, temp)
	return n
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
