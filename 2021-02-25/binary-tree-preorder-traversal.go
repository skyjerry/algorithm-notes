// Leetcode 144. 二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	ans := make([]int, 0)

	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return ans
}