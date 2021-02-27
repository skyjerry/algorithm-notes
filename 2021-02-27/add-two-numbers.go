// Leetcode 2. 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry == 1 {
		val1, val2 := 0, 0

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry
		if sum >= 10 {
			sum = sum - 10
			carry = 1
		} else {
			carry = 0
		}

		head.Next = &ListNode{Val: sum}
		head = head.Next
	}

	return dummy.Next
}