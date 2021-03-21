// Leetcode 103. 二叉树的锯齿形层序遍历

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    queue := make([]*TreeNode, 0)
    neg := true
    ans := make([][]int, 0)

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

        if (!neg) {
            for i := 0; i < n / 2; i++ {
                curLevel[i], curLevel[n - i - 1] = curLevel[n - i - 1], curLevel[i]
            }
        }

        ans = append(ans, curLevel)
        neg = !neg
    }

    return ans

}