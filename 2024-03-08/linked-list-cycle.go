package _024_03_08

/**
206. 反转链表
https://leetcode.cn/problems/reverse-linked-list/

Space: O(1), Time: O(N)
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast, slow := head, head
	for fast.Next == nil && fast != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}
