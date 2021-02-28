// Leetcode 234. 回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	fast, slow := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	mid := slow.Next
	slow.Next = nil

	var prev *ListNode

	for mid != nil {
		next := mid.Next
		mid.Next = prev
		prev = mid
		mid = next
	}

	mid = prev

	for head != nil && mid != nil {
		if head.Val != mid.Val {
			return false
		}

		head = head.Next
		mid = mid.Next
	}

	return true
}