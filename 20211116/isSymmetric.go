package _0211116

/**
 * 对称的二叉树: https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/
 */

func isSymmetric(r *TreeNode) bool {
	if r == nil {
		return true
	}
	return dfs(r.Left, r.Right)
}

func dfs(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	return a.Val == b.Val && dfs(a.Left, b.Right) && dfs(a.Right, b.Left)
}

