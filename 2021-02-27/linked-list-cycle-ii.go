// Leetcode 142. 环形链表 II
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == nil || slow == nil {
			return nil
		}

		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}

	return nil
}