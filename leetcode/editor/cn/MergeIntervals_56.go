package main

import "sort"

//给出一个区间的集合，请合并所有重叠的区间。
//
//
//
// 示例 1:
//
// 输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
//
// 示例 2:
//
// 输入: intervals = [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
// 注意：输入类型已于2019年4月15日更改。 请重置默认代码定义以获取新方法签名。
//
//
//
// 提示：
//
//
// intervals[i][0] <= intervals[i][1]
//
// Related Topics 排序 数组
// 👍 727 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func merge1(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[j][1]
	})
	var res [][]int
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if cur[0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = maxMerge(res[len(res)-1][1], cur[1])
		} else {
			res = append(res, cur)
		}
	}
	return res
}

func maxMerge(x, y int) int {
	if x < y {
		return y
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
