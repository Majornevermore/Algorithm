package main

import "sort"

//给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选
//择一个符号添加在前面。
//
// 返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
//
//
//
// 示例：
//
// 输入：nums: [1, 1, 1, 1, 1], S: 3
//输出：5
//解释：
//
//-1+1+1+1+1 = 3
//+1-1+1+1+1 = 3
//+1+1-1+1+1 = 3
//+1+1+1-1+1 = 3
//+1+1+1+1-1 = 3
//
//一共有5种方法让最终目标和为3。
//
//
//
//
// 提示：
//
//
// 数组非空，且长度不会超过 20 。
// 初始的数组的和不会超过 1000 。
// 保证返回的最终结果能被 32 位整数存下。
//
// Related Topics 深度优先搜索 动态规划
// 👍 504 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(num []int, target int) [][]int {
	var res [][]int
	var dfs func(int, int, []int)
	sort.Ints(num)
	dfs = func(diff int, start int, sub []int) {
		if diff == 0 {
			temp := make([]int, len(sub))
			copy(temp, sub)
			res = append(res, temp)
			return
		}
		per := -1
+1			if num[i] <= diff {
				if num[i] == per {
					continue
				}
				per = num[i]
				sub = append(sub, num[i])
				dfs(diff-num[i], i, sub)
				sub = sub[:len(sub)-1]
			}

		}
	}
	dfs(target, 0, []int{})
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	combinationSum2([]int{1, 1}, 2)
}
