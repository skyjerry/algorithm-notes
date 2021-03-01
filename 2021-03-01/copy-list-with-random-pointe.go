// Leetcode 138. 复制带随机指针的链表
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	node := head
	for node != nil {
		newNode := &Node{
			Val: node.Val,
			Next: node.Next,
		}
		node.Next = newNode
		node = node.Next.Next
	}

	node = head

	for node != nil {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
		node = node.Next.Next
	}

	node = head
	newHead := node.Next

	for node != nil && node.Next != nil {
		tmp := node.Next
		node.Next = node.Next.Next
		node = tmp
	}

	return newHead
}