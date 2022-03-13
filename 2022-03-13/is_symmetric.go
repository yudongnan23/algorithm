package _022_03_13

// 判断对称二叉树
// leetcode链接：https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/

/**
 * 思路：递归
 */

func isSymmetric(r *TreeNode) bool {
	if r == nil {
		return true
	}
	return dfsIII(r.Left, r.Right)
}

func dfsIII(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if r1 == nil {
		return false
	}

	if r2 == nil {
		return false
	}

	if r1.Val != r2.Val {
		return false
	}

	return dfsIII(r1.Left, r2.Right) && dfsIII(r1.Right, r2.Left)
}
