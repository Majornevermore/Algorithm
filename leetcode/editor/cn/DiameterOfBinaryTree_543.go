package main

import "math"

//给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
//
//
//
// 示例 :
//给定二叉树
//
//           1
//         / \
//        2   3
//       / \
//      4   5
//
//
// 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
//
//
//
// 注意：两结点之间的路径长度是以它们之间边的数目表示。
// Related Topics 树
// 👍 575 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	var m int
	if root == nil {
		return 0
	}
	treeDepth(root, &m)
	return m
}

func treeDepth(root *TreeNode, m *int) int {
	if root == nil {
		return 0
	}
	left := treeDepth(root.Left, m)
	right := treeDepth(root.Right, m)
	*m = int(math.Max(float64(left+right), float64(*m)))
	return int(1 + math.Max(float64(left), float64(right)))
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
