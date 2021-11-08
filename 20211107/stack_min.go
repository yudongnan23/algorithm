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