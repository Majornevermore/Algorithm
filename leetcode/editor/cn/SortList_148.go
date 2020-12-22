package main

//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
// 进阶：
//
//
// 你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
//
//
//
//
// 示例 1：
//
//
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
//
//
// 示例 2：
//
//
//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 5 * 104] 内
// -105 <= Node.val <= 105
//
// Related Topics 排序 链表
// 👍 930 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head *ListNode, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return Merge2Node(sort(head, mid), sort(mid, tail))
}

func Merge2Node(head1 *ListNode, head2 *ListNode) *ListNode {
	var new *ListNode
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	if head1.Val < head2.Val {
		new = head1
		new.Next = Merge2Node(head1.Next, head2)
	} else {
		new = head2
		new.Next = Merge2Node(head1, head2.Next)
	}
	return new
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
