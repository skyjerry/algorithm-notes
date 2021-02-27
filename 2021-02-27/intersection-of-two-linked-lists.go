// Leetcode 160. 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	listA, listB := headA, headB

	for listA != listB {
		if listA == nil {
			listA = headB
		} else {
			listA = listA.Next
		}

		if listB == nil {
			listB = headA
		} else {
			listB = listB.Next
		}
	}

	return listA
}