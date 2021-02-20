
// Leetcode 103. Binary Tree Zigzag Level Order Traversal
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	// Init queue, reverse symbol, answer.
	queue := make([]*TreeNode, 0)
	neg := true
	ans := make([][]int, 0)

	// Push root to queue.
	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)
		curLevel := make([]int, n)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			curLevel[i] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// If it should be reversed, then reverse.
		if !neg {
			for i := 0; i < n/2; i++ {
				curLevel[i], curLevel[n-i-1] = curLevel[n-i-1], curLevel[i]
			}
		}

		ans = append(ans, curLevel)
		neg = !neg
	}

	return ans
}