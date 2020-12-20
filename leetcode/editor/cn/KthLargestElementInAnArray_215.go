package main

import (
	"math/rand"
	"time"
)

//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 示例 1:
//
// 输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//
//
// 示例 2:
//
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
//
// 说明:
//
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
// Related Topics 堆 分治算法
// 👍 828 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findKth(a []int, n int, K int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(a, 0, len(a)-1, len(a)-k)
}

func quickSelect(a []int, l, r, index int) int {
	q := RandomPartition(a, l, r)
	if q == index {
		return a[q]
	} else if q > index {
		quickSelect(a, l, q-1, index)
	} else {
		quickSelect(a, q+1, r, index)
	}
}

func RandomPartition(a []int, l, r int) int {
	i := rand.Int()%(r-l+1) + 1
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

func partition(a []int, l, r int) int {
	x := a[r]
	i := l - 1
	for j := 1; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
