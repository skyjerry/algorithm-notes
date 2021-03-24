// Leetcode 105. 从前序与中序遍历序列构造二叉树

func buildTree(preorder []int, inorder []int) *TreeNode {
	return helper(0, 0, len(inorder)-1, preorder, inorder)
}

// build tree
func helper(preStart, inStart, inEnd int, preorder []int, inorder []int) *TreeNode {
	if preStart >= len(preorder) || inStart > inEnd {
		return nil
	}

	root := &TreeNode{Val: preorder[preStart]}
	inIndex := 0

	// find inIndex from inorder
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == root.Val {
			inIndex = i
		}
	}

	// build left tree
	root.Left = helper(preStart+1, inStart, inIndex-1, preorder, inorder)
	// build right tree
	root.Right = helper(preStart+inIndex-inStart+1, inIndex+1, inEnd, preorder, inorder)

	return root
}