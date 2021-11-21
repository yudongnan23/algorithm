package _0211121

/**
 * 删除链表的节点：https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(r *ListNode, v int) *ListNode {
	if r == nil {
		return nil
	}

	p := r
	q := p.Next

	if p.Val == v {
		return q
	}

	for q != nil {
		if q.Val == v {
			p.Next = q.Next
			break
		}
		p = p.Next
		q = q.Next
	}
	return r
}
