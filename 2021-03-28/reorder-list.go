// 143. 重排链表

func reorderList(head *ListNode)  {
    if head == nil {
        return
    }

    preMid := findMid(head)
    secondHead := preMid.Next
    preMid.Next = nil

    secondHead = reverse(secondHead)

    var next *ListNode
    for secondHead != nil {
        next = secondHead.Next
        secondHead.Next = head.Next
        head.Next = secondHead
        head = secondHead.Next
        secondHead = next
    }
}

func findMid(head *ListNode) *ListNode {
    slow, fast := head, head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    return slow
}

func reverse(head *ListNode) *ListNode {
    var prev, next *ListNode
    for head != nil {
        next = head.Next
        head.Next = prev
        prev = head
        head = next
    }

    return prev
}
