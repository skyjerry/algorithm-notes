// Leetcode 83. 删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := head

	for newHead != nil && newHead.Next != nil {
		if newHead.Val == newHead.Next.Val {
			newHead.Next = newHead.Next.Next
		} else {
			newHead = newHead.Next
		}
	}

	return head
}