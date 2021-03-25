// Leetcode 19. 删除链表的倒数第 N 个结点

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Val: 0, Next: head}

    fast, slow := dummy, dummy

    for i := 0; i <= n; i++ {
        fast = fast.Next
    }

    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }

    slow.Next = slow.Next.Next

    return dummy.Next
}