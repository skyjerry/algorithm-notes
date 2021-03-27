// Leetcode 82. 删除排序链表中的重复元素 II

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	head = dummy

	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			repeatVal := head.Next.Val
			for head.Next != nil && repeatVal == head.Next.Val {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}

	return dummy.Next
}