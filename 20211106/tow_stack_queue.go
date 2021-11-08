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
