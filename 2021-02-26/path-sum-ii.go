// Leetcode 113. 路径总和 II
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	helper(root, []int{}, sum, &result)
	return result
}

func helper(root *TreeNode, path []int, sum int, result *[][]int) {
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			*result = append(*result, append(path, root.Val))
		}
		return
	}
	if root.Left != nil {
		helper(root.Left, append(append([]int{}, path...), root.Val), sum-root.Val, result)
	}
	if root.Right != nil {
		helper(root.Right, append(append([]int{}, path...), root.Val), sum-root.Val, result)
	}
}