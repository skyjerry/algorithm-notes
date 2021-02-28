// Leetcode 23. 合并K个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}

	if n == 1 {
		return lists[0]
	}

	if n == 2 {
		return mergeLists(lists[0], lists[1])
	}

	return mergeLists(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))
}

func mergeLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	dummy := &ListNode{Val: 0}
	head := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}

		head = head.Next
	}

	if l1 != nil {
		head.Next = l1
	}

	if l2 != nil {
		head.Next = l2
	}

	return dummy.Next
}