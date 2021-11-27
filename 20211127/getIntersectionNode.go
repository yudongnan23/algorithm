package _0211127

/**
 * 两个链表的公共节点：https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(a, b *ListNode) *ListNode {
	dynamicA := a
	dynamicB := b

	for dynamicA != dynamicB {
		if dynamicA == nil {
			dynamicA = b
		} else {
			dynamicA = dynamicA.Next
		}

		if dynamicB == nil {
			dynamicB = a
		} else {
			dynamicB = dynamicB.Next
		}
	}
	return dynamicA
}
