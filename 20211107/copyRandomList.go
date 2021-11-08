package _0211107

/**
 * 复杂链表的复制： https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	nodeMapping := make(map[*Node]*Node, 0)
	randomMapping := make(map[*Node]*Node, 0)

	h := &Node{Val: head.Val}
	point := h

	nodeMapping[head] = h
	if head.Random != nil {
		randomMapping[h] = head.Random
	}

	cur := head.Next
	for cur != nil {
		n := &Node{Val: cur.Val}
		if cur.Random != nil {
			randomMapping[n] = cur.Random
		}

		nodeMapping[cur] = n
		point.Next = n
		point = point.Next

		cur = cur.Next
	}

	for r, v := range randomMapping {
		r.Random = nodeMapping[v]
	}

	return h
}
