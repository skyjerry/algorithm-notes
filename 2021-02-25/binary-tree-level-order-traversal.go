// Leetcode 102. 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
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

	return ans
}