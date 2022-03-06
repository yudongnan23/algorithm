package _022_03_06

// 从上大小打印二叉树II
// leetcode链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/

// 思路：广度遍历

type TreeNodeWithDepth struct {
	node  *TreeNode
	depth int
}

func levelOrderII(r *TreeNode) [][]int {
	queue := []*TreeNodeWithDepth{{node: r, depth: 0}}
	res := make([][]int, 0)

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.node == nil {
			break
		}

		if len(res) == cur.depth {
			res = append(res, make([]int, 0))
		}

		res[cur.depth] = append(res[cur.depth], cur.node.Val)

		if cur.node.Left != nil {
			queue = append(queue, &TreeNodeWithDepth{
				node:  cur.node.Left,
				depth: cur.depth + 1,
			})
		}

		if cur.node.Right != nil {
			queue = append(queue, &TreeNodeWithDepth{
				node:  cur.node.Right,
				depth: cur.depth + 1,
			})
		}
	}

	return res
}
