// Leetcode 110. 平衡二叉树
func isBalanced(root *TreeNode) bool {
	depth := maxDepth(root)
	if depth == -1 {
		return false
	}

	return true
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}

	if left > right {
		return left + 1
	}

	return right + 1
}