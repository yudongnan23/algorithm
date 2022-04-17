package _022_04_17

// 删除链表的节点
// leetcode链接：https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/

/**
 * 思路：
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(h *ListNode, val int) *ListNode {
	if h != nil && h.Val == val {
		return h.Next
	}

	p := &ListNode{
		Val:  0,
		Next: h,
	}

	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
			break
		}

		p = p.Next
	}

	return h
}
