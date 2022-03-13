package _022_03_13

// 树的子结构
// leetcode链接：https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/

/**
 * 思路：先对二叉树进行遍历，然后逐个进行子树判断
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	queue := []*TreeNode{A}

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.Val == B.Val && dfs(cur, B) {
			return true
		}

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}

	}
	return false
}

func dfs(A *TreeNode, B *TreeNode) bool {
	if A == nil {
		return false
	}

	if A.Val == B.Val {
		if B.Left != nil && B.Right != nil {
			return dfs(A.Left, B.Left) && dfs(A.Right, B.Right)
		}
		if B.Left != nil {
			return dfs(A.Left, B.Left)
		}
		if B.Right != nil {
			return dfs(A.Right, B.Right)
		}
		return true
	}

	return false
}
