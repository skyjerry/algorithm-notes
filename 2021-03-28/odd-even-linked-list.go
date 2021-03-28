// 328. 奇偶链表

func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    odd := head
    even := head.Next
    evenHead := even

    for even != nil && even.Next != nil {
        odd.Next = odd.Next.Next
        even.Next = even.Next.Next
        odd = odd.Next
        even = even.Next
    }
    odd.Next = evenHead

    return head
}