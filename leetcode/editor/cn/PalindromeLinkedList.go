package main

//请判断一个链表是否为回文链表。
//
// 示例 1:
//
// 输入: 1->2
//输出: false
//
// 示例 2:
//
// 输入: 1->2->2->1
//输出: true
//
//
// 进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
// Related Topics 链表 双指针
// 👍 791 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var Left *ListNode

func isPalindrome(head *ListNode) bool {
	Left = head
	return traverse(head)
}

// 反正树的遍历方法，前序遍历和后续遍历得到的值应该是一样的
func traverse(right *ListNode) bool {
	if right == nil {
		return true
	}
	res := traverse(right.Next)
	res = res && (right.Val == Left.Val)
	Left = Left.Next
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
