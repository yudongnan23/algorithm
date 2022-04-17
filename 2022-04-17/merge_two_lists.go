package _022_04_17

// 合并两个排序链表
// leetcode链接：https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/

/**
 * 思路：双指针
 */

const (
	_first int = iota
	_second
	_firstNil
	_secondNil
)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var h *ListNode
	var p *ListNode

	for l1 != nil || l2 != nil {
		val, opt := compareNode(l1, l2)
		curNode := &ListNode{
			Val: val,
		}

		if h == nil {
			h = curNode
			p = curNode
		} else {
			p.Next = curNode
			p = p.Next
		}

		if opt == _firstNil {
			p.Next = l2.Next
			break
		}
		if opt == _secondNil {
			p.Next = l1.Next
			break
		}

		if opt == _first {
			l1 = l1.Next
		}

		if opt == _second {
			l2 = l2.Next
		}
	}

	return h
}

func compareNode(l1, l2 *ListNode) (int, int) {
	if l1 == nil {
		return l2.Val, _firstNil
	}
	if l2 == nil {
		return l1.Val, _secondNil
	}
	if l1.Val < l2.Val {
		return l1.Val, _first
	}
	return l2.Val, _second
}
