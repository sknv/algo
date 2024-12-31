package leetcode

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/

func lowestCommonAncestor(root, p, q *IntTreeNode) *IntTreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}
