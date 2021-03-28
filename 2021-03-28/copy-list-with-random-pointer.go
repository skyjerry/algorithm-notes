// Leetcode 138. 复制带随机指针的链表

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }

    oldHead := head
    
    for head != nil {
        head.Next = &Node{
            Val: head.Val,
            Next: head.Next,
        }
        head = head.Next.Next
    }

    head = oldHead
    
    for head != nil {
        if head.Random != nil {
            head.Next.Random = head.Random.Next
        }
        head = head.Next.Next
    }

    head = oldHead
    newHead := head.Next

    for head != nil && head.Next != nil {
        next := head.Next
        head.Next = head.Next.Next
        head = next
    }

    return newHead
}