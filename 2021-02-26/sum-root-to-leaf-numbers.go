// Leetcode 129. 求根到叶子节点数字之和
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return helper(root, 0)
}

func helper(root *TreeNode, currSum int) int {
	if root == nil {
		return 0
	}

	currSum = currSum*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return currSum
	}

	return helper(root.Left, currSum) + helper(root.Right, currSum)
}