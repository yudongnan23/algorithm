package _0211121

/**
 * 合并两个排序链表：https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var h *ListNode
	p := l1
	q := l2
	if l2.Val < l1.Val {
		h = l2
		q = q.Next
	} else {
		h = l1
		p = p.Next
	}
	c := h

	for p != nil || q != nil {
		if p == nil {
			c.Next = q
			break
		}
		if q == nil {
			c.Next = p
			break
		}

		if p.Val < q.Val {
			c.Next = p
			p = p.Next
		} else {
			c.Next = q
			q = q.Next
		}

		c = c.Next
	}

	return h
}

func (l *ListNode) Value() []int {
	v := make([]int, 0)
	for l != nil {
		v = append(v, l.Val)
	}
	return v
}
