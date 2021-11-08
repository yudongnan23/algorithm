package _0211107

/**
 * 反转链表：https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	pre := head
	cur := head.Next
	pre.Next = nil
	for {
		if cur == nil {
			break
		}

		n := cur.Next
		cur.Next = pre
		pre = cur
		cur = n
	}

	return pre
}

func (l *ListNode)getValueAsList()[]int {
	rlt := make([]int, 0)
	for l != nil {
		rlt = append(rlt, l.Val)
		l = l.Next
	}
	return rlt
}