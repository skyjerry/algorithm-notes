// Leetcode 107. 二叉树的层序遍历 II
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	var ans [][]int

	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)
		level := make([]int, 0)

		for n > 0 {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			n--
		}
		ans = append(ans, level)
	}

	n := len(ans)
	for i := 0; i < n/2; i++ {
		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
	}

	return ans
}