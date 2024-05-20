package binarytreedfs

func pathSum(root *TreeNode, targetSum int) int {
	diffMap := make(map[int]int)
	diffMap[0] = 1

	var ans int
	dfs(root, targetSum, 0, diffMap, &ans)
	return ans
}

func dfs(node *TreeNode, targetSum int, sum int, diffMap map[int]int, ans *int) {
	if node == nil {
		return
	}

	sum += node.Val
	diff := sum - targetSum
	count, ok := diffMap[diff]
	if ok {
		*ans += count
	}
	diffMap[sum]++

	if node.Left != nil {
		dfs(node.Left, targetSum, sum, diffMap, ans)
	}

	if node.Right != nil {
		dfs(node.Right, targetSum, sum, diffMap, ans)
	}

	diffMap[sum]--
}
