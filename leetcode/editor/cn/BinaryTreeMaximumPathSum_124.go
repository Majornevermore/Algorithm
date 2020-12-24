package main

import "math"

//给定一个非空二叉树，返回其最大路径和。
//
// 本题中，路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
//
//
//
// 示例 1：
//
// 输入：[1,2,3]
//
//       1
//      / \
//     2   3
//
//输出：6
//
//
// 示例 2：
//
// 输入：[-10,9,20,null,null,15,7]
//
//   -10
//   / \
//  9  20
//    /  \
//   15   7
//
//输出：42
// Related Topics 树 深度优先搜索 递归
// 👍 824 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//func maxPathSum(root *TreeNode) int {
//	maxSum := math.MinInt32
//	var dfs func(*TreeNode) int
//	dfs = func(node *TreeNode) int {
//		if node == nil {
//			return 0
//		}
//		left := maxPath(dfs(node.Left), 0)
//		right := maxPath(dfs(node.Right), 0)
//		temp := maxPath(left, right) + node.Val
//		// 节点最大路径取决于该节点的值和该节点的左右自己子节点的最大贡献值
//		priceValue := left + right + root.Val
//		maxSum = maxPath(priceValue, maxSum)
//		// 节点最大贡献值
//		return temp
//	}
//	dfs(root)
//	return maxSum
//}
//
//func maxPath(x, y int) int {
//	if x < y {
//		return y
//	}
//	return x
//}
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右子节点的最大贡献值
		// 只有在最大贡献值大于 0 时，才会选取对应子节点
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
		priceNewPath := node.Val + leftGain + rightGain

		// 更新答案
		maxSum = max(maxSum, priceNewPath)

		// 返回节点的最大贡献值
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
