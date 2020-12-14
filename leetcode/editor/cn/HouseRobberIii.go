package main

//在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“
//房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
//
// 计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。
//
// 示例 1:
//
// 输入: [3,2,3,null,3,null,1]
//
//     3
//    / \
//   2   3
//    \   \
//     3   1
//
//输出: 7
//解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
//
// 示例 2:
//
// 输入: [3,4,5,1,3,null,1]
//
//     3
//    / \
//   4   5
//  / \   \
// 1   3   1
//
//输出: 9
//解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.
//
// Related Topics 树 深度优先搜索
// 👍 664 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var hashmapTree = make(map[*TreeNode]int)

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if v, ok := hashmapTree[root]; ok {
		return v
	}
	var rob_yes int
	var rob_not int
	var rob_left int
	var rob_right int
	if root.Left != nil {
		rob_left = rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		rob_right = rob(root.Right.Left) + rob(root.Right.Right)
	}
	rob_yes = rob_left + rob_right + root.Val
	rob_not = rob(root.Right) + rob(root.Left)
	value := maxRobb(rob_yes, rob_not)
	hashmapTree[root] = value
	return value
}

func maxRobb(x, y int) int {
	if x < y {
		return y
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
