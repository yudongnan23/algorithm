package _0211107

/**
 * 包含min函数的栈：https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/
 */

type MinStack struct {
	stack  []int
	min    int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (c *MinStack) Push(x int) {
	c.stack = append(c.stack, x)
	if x < c.min || len(c.stack) == 1 {
		c.min = x
	}
	return
}

func (c *MinStack) Pop() {
	if len(c.stack) == 0 {
		return
	}

	c.stack = c.stack[0 : len(c.stack)-1]
	c.resetMin()
	return
}

func (c *MinStack) resetMin() {
	min := 0

	if len(c.stack) != 0 {
		min = c.stack[0]
	}
	for i := 0; i < len(c.stack); i++ {
		if c.stack[i] < min {
			min = c.stack[i]
		}
	}
	c.min = min
	return
}

func (c *MinStack) Top() int {
	if len(c.stack) == 0 {
		return 0
	}
	return c.stack[len(c.stack)-1]
}

func (c *MinStack) Min() int {
	return c.min
}

// 以下为双向链表实现
//type MinStack struct {
//	head *node
//	tail *node
//}
//
//type node struct {
//	v       int
//	next    *node
//	pre     *node
//	curMin  int
//}
//
///** initialize your data structure here. */
//func Constructor() MinStack {
//	m := MinStack{
//		head:  nil,
//		tail: nil,
//	}
//	return m
//}
//
//func (m *MinStack) Push(v int)  {
//	n := node{
//		v: v,
//	}
//
//	if m.head == nil {
//		n.curMin = v
//		m.head = &n
//		m.tail = &n
//		return
//	}
//
//	n.curMin = m.tail.curMin
//	if v < m.tail.curMin {
//		n.curMin = v
//	}
//
//	n.pre = m.tail
//	m.tail.next = &n
//	m.tail = m.tail.next
//
//	return
//}
//
//func (m *MinStack) Pop()  {
//	if m.tail == nil {
//		return
//	}
//
//	if m.tail == m.head {
//		m.head = nil
//		m.tail = nil
//		return
//	}
//
//	m.tail = m.tail.pre
//}
//
//func (m *MinStack) Top() int {
//	if m.tail == nil {
//		return 0
//	}
//
//	return m.tail.v
//}
//
//func (m *MinStack) Min() int {
//	return m.tail.curMin
//}
