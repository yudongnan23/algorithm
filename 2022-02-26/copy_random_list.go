package _022_02_26

// 复杂链表的复制
// leecode链接：https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/

/**
 * 思路：使用两个map，一个map存储存在random指针的新节点与老节点的映射，一个map存储老节点与新节点的映射
 * 第一遍遍历，复制完整链表
 * 第二遍遍历，连接random节点
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

type withRandomNode *Node

type oldNode *Node

type newNode *Node

func copyRandomList(h *Node) *Node {
	if h == nil {
		return nil
	}

	randomMapping := make(map[withRandomNode]oldNode, 0)
	oldNewMapping := make(map[oldNode]newNode)

	newH := &Node{Val: h.Val}
	pointFirst := newH

	if h.Random != nil {
		randomMapping[newH] = h.Random
	}
	oldNewMapping[h] = newH

	h = h.Next
	for h != nil {
		curNode := &Node{Val: h.Val}

		if h.Random != nil {
			randomMapping[curNode] = h.Random
		}
		oldNewMapping[h] = curNode

		pointFirst.Next = curNode
		h = h.Next
		pointFirst = pointFirst.Next
	}

	pointSecond := newH

	for pointSecond != nil {
		if old := randomMapping[pointSecond]; old != nil {
			pointSecond.Random = oldNewMapping[old]
		}

		pointSecond = pointSecond.Next
	}

	return newH
}
