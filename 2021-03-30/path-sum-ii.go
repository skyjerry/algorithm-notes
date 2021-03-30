// Leetcode 113. 路径总和 II

func pathSum(root *TreeNode, targetSum int) [][]int {
    if root == nil {
        return nil
    }

    ans := make([][]int, 0)
    helper(root, []int{}, targetSum, &ans)

    return ans
}

func helper(root *TreeNode, path []int, sum int, ans *[][]int) {
    if root.Left == nil && root.Right == nil {
        if root.Val == sum {
            *ans = append(*ans, append(path, root.Val))
        }

        return
    }

    if root.Left != nil {
        helper(root.Left, append(append([]int{}, path...), root.Val), sum - root.Val, ans)
    }
    if root.Right != nil {
        helper(root.Right, append(append([]int{}, path...), root.Val), sum - root.Val, ans)
    }
}
