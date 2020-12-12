package main

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例：
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
// Related Topics 链表
// 👍 1433 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var mergeHead *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		mergeHead = l1
		mergeHead.Next = mergeTwoLists(l1.Next, l2)
	} else {
		mergeHead = l2
		mergeHead.Next = mergeTwoLists(l1, l2.Next)
	}
	return mergeHead
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
