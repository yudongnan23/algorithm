package _022_02_26

// 反转链表
// leetcode链接：https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
// 思路：遍历链表，将链表进行反向连接

func reverseList(h *ListNode) *ListNode {
	var newH *ListNode

	for h != nil {
		hNext := h.Next

		h.Next = newH
		newH = h

		h = hNext
	}

	return newH
}
