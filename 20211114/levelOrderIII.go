package _0211114

/**
 * 从上到下打印二叉树III：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
 */

func levelOrderIII(r *TreeNode) [][]int {
	rlt := make([][]int, 0)
	if r == nil {
		return rlt
	}

	q := make([]*TreeNode, 1)
	dump := make(map[*TreeNode]int, 0)
	q[0] = r
	dump[r] = 0

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		if len(rlt) <= dump[cur] {
			rlt = append(rlt, []int{cur.Val})
		} else {
			if dump[cur]%2 == 0 {
				rlt[dump[cur]] = append(rlt[dump[cur]], cur.Val)
			} else {
				rlt[dump[cur]] = append([]int{cur.Val}, rlt[dump[cur]]...)
			}

		}

		if cur.Left != nil {
			dump[cur.Left] = dump[cur] + 1
			q = append(q, cur.Left)
		}

		if cur.Right != nil {
			dump[cur.Right] = dump[cur] + 1
			q = append(q, cur.Right)
		}
	}
	return rlt
}
