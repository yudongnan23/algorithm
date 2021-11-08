package _0211107

/**
 * 从头到尾打印链表：https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/submissions/
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	rlt := make([]int, 10000)
	i := 9999
	for head != nil{
		rlt[i] = head.Val
		head = head.Next
		i --
	}

	return rlt[i+1:]
}
