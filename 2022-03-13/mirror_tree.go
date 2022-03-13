package _022_03_13

// 树的镜像
// leetcode链接：https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/

/**
 * 思路：递归
 */

func mirrorTree(r *TreeNode) *TreeNode {
	dfsII(r)
	return r
}

func dfsII(r *TreeNode) {
	if r == nil {
		return
	}
	left := r.Left
	r.Left = r.Right
	r.Right = left

	dfsII(r.Left)
	dfsII(r.Right)
}
