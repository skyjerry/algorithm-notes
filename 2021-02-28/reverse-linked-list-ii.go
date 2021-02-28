// Leetcode 92. 反转链表 II
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Val: 0}
	dummy.Next = head

	firstTail := dummy
	i := 0
	for i < left-1 {
		firstTail = firstTail.Next
		i++
	}

	secHead := firstTail.Next
	newTail := secHead
	firstTail.Next = nil

	var prev *ListNode
	for i := 0; i <= right-left; i++ {
		next := secHead.Next
		secHead.Next = prev
		prev = secHead
		secHead = next
	}

	newTail.Next = secHead
	firstTail.Next = prev

	return dummy.Next
}