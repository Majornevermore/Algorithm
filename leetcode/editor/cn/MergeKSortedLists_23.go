package main

//给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
//
// 示例 2：
//
// 输入：lists = []
//输出：[]
//
//
// 示例 3：
//
// 输入：lists = [[]]
//输出：[]
//
//
//
//
// 提示：
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
// Related Topics 堆 链表 分治算法
// 👍 1049 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeList(lists, 0, len(lists)-1)
}

func mergeList(list []*ListNode, l, r int) *ListNode {
	mid := (l + r) / 2
	if l == r {
		return list[l]
	}
	u := mergeList(list, l, mid)
	v := mergeList(list, mid+1, r)
	return mergeTwoList(u, v)
}

func mergeTwoList(a *ListNode, b *ListNode) *ListNode {
	var new *ListNode
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val < b.Val {
		new = a
		new.Next = mergeTwoList(a.Next, b)
	} else {
		new = b
		new.Next = mergeTwoList(a, b.Next)
	}
	return new
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
