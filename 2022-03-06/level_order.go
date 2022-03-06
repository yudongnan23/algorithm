package _022_03_06

// 从上到下打印二叉树
// leetcode链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/

// 思路：

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(r *TreeNode) []int {
	queue := []*TreeNode{r}
	res := make([]int, 0)

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur == nil {
			break
		}

		res = append(res, cur.Val)

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}

		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

	return res
}
