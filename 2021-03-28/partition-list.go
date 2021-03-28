// Leetcode 86. 分隔链表

func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return nil
    }

    bigDummy := &ListNode{Next: head}
    smallDummy := &ListNode{Next: nil}

    bigHead := bigDummy
    smallHead := smallDummy

    for bigHead.Next != nil {
        if bigHead.Next.Val < x {
            smallHead.Next = bigHead.Next
            smallHead = smallHead.Next
            bigHead.Next = bigHead.Next.Next
        } else {
            bigHead = bigHead.Next
        }
    }

    smallHead.Next = bigDummy.Next

    return smallDummy.Next
}