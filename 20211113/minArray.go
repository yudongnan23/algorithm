package _0211113

/**
 * 旋转数组中的最小数：https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
 */

func minArray(n []int) int {
	if len(n) == 0 {
		return 0
	}

	l := 0
	r := len(n) - 1

	for l < r {
		m := (l + r) / 2

		if n[m] < n[l] {
			r = m
		} else if n[m] > n[r] {
			l = m + 1
		}else {
			r --
		}
	}

	return n[l]
}
