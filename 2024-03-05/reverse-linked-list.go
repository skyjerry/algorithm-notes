// Leetcode 206. 反转链表

// Space: O(1), Time: O(N)

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