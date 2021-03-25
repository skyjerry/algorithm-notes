// Leetcode 142. 环形链表 II

func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }


    slow, fast := head, head
    
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            slow = head
            for slow != fast {
                slow = slow.Next
                fast = fast.Next
            }

            return fast
        }
    }

    return nil
}