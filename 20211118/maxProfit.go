package _0211118

/**
 * 股票的最大利润： https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/
 */

func maxProfit(p []int) int {
	if len(p) == 0 {
		return 0
	}

	min := p[0]
	max := -p[0]
	for i := 1; i < len(p); i++ {
		if p[i]-min > max {
			max = p[i] - min
		}

		if p[i] < min {
			min = p[i]
		}
	}
	if max <= 0 {
		max = 0
	}
	return max
}
