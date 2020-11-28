package main

//给定一个二叉树，找出其最大深度。
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回它的最大深度 3 。
// Related Topics 树 深度优先搜索 递归
// 👍 747 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{}
	var res int
	q = append(q, root)
	for len(q) > 0 {
		leverSize := len(q)
		for i := 0; i < leverSize; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		res++
	}
	return res
}

func maxDepth11(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + maxDepth11(maxDepth(root.Left), maxDepth(root.Right))
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
