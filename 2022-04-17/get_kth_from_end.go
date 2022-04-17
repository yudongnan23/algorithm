package _022_04_17

// 链表中倒数第k个节点
// leetcode链接：https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

/**
 * 思路：双指针
 */

func getKthFromEnd(h *ListNode, k int) *ListNode {
	r := h
	for i := 0; r != nil && i < k; i++ {
		r = r.Next
	}
	if r == nil {
		return nil
	}

	l := h
	for r != nil {
		r = r.Next
		l = l.Next
	}

	return l
}
