package binarytreedfs

func maxDepth(root *TreeNode) int {
	r := Res{}
	r.dfs(root, 0)
	return r.Depth
}

type Res struct {
	Depth int
}

func (r *Res) dfs(node *TreeNode, count int) {

	if node != nil {
		r.dfs(node.Left, count+1)
		r.dfs(node.Right, count+1)
	}
	if count > r.Depth {
		r.Depth = count
	}
}
