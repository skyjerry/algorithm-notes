// Leetcode 86. 分隔链表
func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return head
    }
    smallDummy := &ListNode{
        Val: 0,
        Next: nil,
    }
    bigDummy := &ListNode{
        Val: 0,
        Next: head,
    }

    smallList := smallDummy
    bigList := bigDummy

    for bigList.Next != nil {
        if bigList.Next.Val < x {
            smallList.Next = bigList.Next
            smallList = smallList.Next
            bigList.Next = bigList.Next.Next
        } else {
            bigList = bigList.Next
        }
    }

    smallList.Next = bigDummy.Next
    return smallDummy.Next

}