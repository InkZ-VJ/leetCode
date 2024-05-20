package binarytreedfs

func goodNodes(root *TreeNode) int {
	return goodNodeCheck(root, -1e4)
}

func goodNodeCheck(node *TreeNode, x int) int {
	if node == nil {
		return 0
	}

	x, checker := max(node.Val, x)
	count := 0
	if checker {
		count = 1
	}
	return count + goodNodeCheck(node.Left, x) + goodNodeCheck(node.Right, x)
}

func max(a, b int) (int, bool) {
	if a >= b {
		return a, true
	}

	return b, false
}
