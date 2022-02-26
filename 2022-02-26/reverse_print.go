package _022_02_26

// 从尾到头打印链表
// leetcode链接：https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
// 思路：遍历链表，逆序放入切片中

type ListNode struct {
	V    int
	Next *ListNode
}

const _maxCount = 10000

func reversePrint(h *ListNode) []int {
	res := make([]int, _maxCount)

	i := 9999

	for h != nil {
		res[i] = h.V

		h = h.Next

		i--
	}

	return res[i+1:]
}
