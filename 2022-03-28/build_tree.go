package _022_03_27

// 重建二叉树
// leetcode链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/

/**
 * 思路：递归
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func findIndex(slice []int, target int) int {
	var res int
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			res = i
			break
		}
	}
	return res
}

func buildTree(pre []int, ino []int) *TreeNode {
	if len(pre) == 0 || len(ino) == 0 {
		return nil
	}

	rootIndex := findIndex(ino, pre[0])

	r := TreeNode{
		Val:   pre[0],
		Left:  buildTree(pre[1:rootIndex+1], ino[:rootIndex]),
		Right: buildTree(pre[rootIndex+1:], ino[rootIndex+1:]),
	}

	return &r
}