// Leetcode 143. 重排链表
func reorderList(head *ListNode)  {
    if head == nil {
        return
    }

    mid := findMid(head)
	tail := mid.Next
	mid.Next = nil
	head2 := reverseList(tail)
    var next *ListNode

    for head2 != nil {
        next = head2.Next
        head2.Next = head.Next
        head.Next = head2
        head = head2.Next
		head2 = next
	}
}

func findMid(head *ListNode) *ListNode {
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }

    return slow
}

func reverseList(head *ListNode) *ListNode {
    var prev, next *ListNode
    for head != nil {
        next = head.Next
        head.Next = prev
        prev = head
        head = next
    }
    return prev
}