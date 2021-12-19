package _0211106

/**
 * 用两个栈实现一个队列：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/submissions/
 */

type CQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (c *CQueue) AppendTail(value int) {
	c.stack1 = append(c.stack1, value)
}

func (c *CQueue) DeleteHead() int {
	if len(c.stack2) == 0 {
		c.transfer()
	}

	if len(c.stack2) == 0 {
		return -1
	}

	value := c.stack2[0]
	c.stack2 = c.stack2[1:]

	return value
}

func (c *CQueue) transfer() {
	c.stack2 = c.stack1
	c.stack1 = []int{}
}

// 链表方式实现
//
//type CQueue struct {
//	l1 *linkList
//	l2 *linkList
//}
//
//type linkList struct {
//	head *node
//	tail *node
//}
//
//type node struct {
//	v    int
//	next *node
//	pre  *node
//}
//
//func Constructor() CQueue {
//	q := CQueue{
//		l1: &linkList{},
//		l2: &linkList{},
//	}
//	return q
//}
//
//// 队列添加元素
//// 只添加在第一个栈
//func (c *CQueue) AppendTail(v int)  {
//	n := &node{
//		v: v,
//	}
//
//	// 栈为空
//	if c.l1.head == nil {
//		c.l1.head = n
//		c.l1.tail = n
//		return
//	}
//
//	n.pre = c.l1.tail
//	c.l1.tail.next = n
//	c.l1.tail = c.l1.tail.next
//	return
//}
//
//func (c *CQueue) DeleteHead() int {
//	c.move()
//
//	if c.l2.head == nil {
//		return -1
//	}
//
//	v := c.l2.head.v
//	c.l2.head = c.l2.head.next
//
//	return v
//}
//
//func (c *CQueue)move() {
//	if c.l2.head != nil {
//		return
//	}
//
//	c.l2.head = c.l1.head
//	c.l2.tail = c.l1.tail
//	c.l1.head = nil
//	c.l1.tail = nil
//
//	return
//}