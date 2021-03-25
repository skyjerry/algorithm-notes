// 2. 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{Val: 0}
    head, carry := dummy, 0

    for l1 != nil || l2 != nil {
        sum := 0
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
		}

        sum += carry
        if sum > 9 {
            carry = 1
            sum = sum - 10
        } else {
            carry = 0
        }
        head.Next = &ListNode{Val: sum}
        head = head.Next
    }
    if carry == 1 {
        head.Next = &ListNode{Val: 1}
    }
	
    return dummy.Next
}