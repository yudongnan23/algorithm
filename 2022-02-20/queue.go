package _022_02_20

// 使用两个栈实现一个队列
// leetcode链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
// 思路：使用双向链表实现

type Queue struct {
	head *DoubleLinkedListNode
	tail *DoubleLinkedListNode
}

func Constructor() Queue {
	return Queue{
		head: nil,
		tail: nil,
	}
}

func (q *Queue) AppendTail(v int) {
	node := DoubleLinkedListNode{
		V: v,
	}

	if q.head == nil {
		q.head = &node
		q.tail = &node
		return
	}

	node.Pre = q.tail
	q.tail.Next = &node
	q.tail = &node
	return
}

func (q *Queue) DeleteHead() int {
	if q.head == nil {
		return -1
	}

	res := q.head.V
	q.head = q.head.Next

	return res
}
