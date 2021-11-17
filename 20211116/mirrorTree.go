package _0211116

/**
 * 二叉树的镜像:https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(r *TreeNode) *TreeNode {
	if r == nil {
		return nil
	}
	t := r.Left
	r.Left = r.Right
	r.Right = t
	mirrorTree(r.Left)
	mirrorTree(r.Right)
	return r
}
