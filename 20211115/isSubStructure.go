package _0211115

/**
 * 树的子结构：https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return false
	}

	q := make([]*TreeNode, 1)
	if a != nil {
		q[0] = a
	}

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		if cur.Val == b.Val && isSub(cur, b) {
			return true
		}

		if cur.Left != nil {
			q = append(q, cur.Left)
		}

		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}
	return false
}

func isSub(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if b == nil {
		return true
	}
	if a.Val != b.Val {
		return false
	}
	return isSub(a.Left, b.Left) && isSub(a.Right, b.Right)
}