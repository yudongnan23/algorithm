package _0211121

/**
 * 链表中倒数第k个节点：https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
 */

func getKthFromEnd(r *ListNode, k int) *ListNode {
	if r == nil {
		return nil
	}
	p := r
	q := r

	// 先让右边指针走k步
	for i := 0; i < k; i++ {
		if q == nil {
			return nil
		}
		q = q.Next
	}

	// 两个指针开始同时走
	for q != nil {
		p = p.Next
		q = q.Next
	}

	return p
}
