// Leetcode 124. 二叉树中的最大路径和
var maxSum int

func maxPathSum(root *TreeNode) int {
	maxSum = -1 << 31
	dfs(root)
	return maxSum
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := max(0, dfs(root.Left))
	right := max(0, dfs(root.Right))
	maxSum = max(maxSum, left+root.Val+right)
	return max(left, right) + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}