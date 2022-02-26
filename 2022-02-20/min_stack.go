package _022_02_20

// 实现包含min函数的栈
// leetcode链接：https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/
// 思路：双向链表实现

type MinStack struct {
	head *node
}

type node struct {
	v      int
	next   *node
	CurMin int
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{
		head: nil,
	}
}

func (m *MinStack) Push(v int) {
	n := node{
		v:      v,
		CurMin: v,
	}

	if m.head == nil {
		m.head = &n
		return
	}

	if m.head.CurMin < v {
		n.CurMin = m.head.CurMin
	}

	n.next = m.head
	m.head = &n
	return
}

func (m *MinStack) Pop() {
	m.head = m.head.next
	return
}

func (m *MinStack) Top() int {
	if m.head == nil {
		return 0
	}
	return m.head.v
}

func (m *MinStack) Min() int {
	return m.head.CurMin
}
