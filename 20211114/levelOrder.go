package _0211114

/**
 * 从上到下打印二叉树：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(r *TreeNode) []int {
	rlt := make([]int, 0)
	if r == nil {
		return rlt
	}

	q := make([]*TreeNode, 1)
	q[0] = r

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		rlt = append(rlt, cur.Val)

		if cur.Left != nil {
			q = append(q, cur.Left)
		}

		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}

	return rlt
}
