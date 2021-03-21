// Leetcode 235. 二叉搜索树的最近公共祖先

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    if root.Val > p.Val && root.Val > q.Val {
        return lowestCommonAncestor(root.Left, p, q)
    } else if root.Val < p.Val && root.Val < q.Val {
        return lowestCommonAncestor(root.Right, p, q)
    } else {
        return root
    }
}