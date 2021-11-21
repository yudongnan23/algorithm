package _0211121

import (
	"math"
	"strconv"
)

/**
 * 把数字翻译成字符串：https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
 */

func translateNum(n int) int {
	l := len(strconv.Itoa(n))
	d := make([]int, l+1)
	d[0] = 1
	d[1] = 1

	for i := 1; i < l; i++ {
		d[i+1] = d[i]
		if andOne(n, l, i) {
			d[i+1] = d[i] + d[i-1]
		}
	}
	return d[len(d)-1]
}

func andOne(n, l, i int) bool {
	t := n / int(math.Pow(10, float64(l-i-1)))
	return t%100 != t%10 && t%100 <= 25
}
