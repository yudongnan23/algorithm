package _022_03_19

// 股票最大利润
// leetcode链接：https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/

/**
 * 思路：动态规划
 */

type profit struct {
	maxProfit int
	minIdx    int
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	dynamic := profit{
		maxProfit: 0,
		minIdx:    0,
	}

	for i := 1; i < len(prices); i++ {
		max := Max(prices[i]-prices[dynamic.minIdx], dynamic.maxProfit)
		if max > dynamic.maxProfit {
			dynamic.maxProfit = max
		}
		max = Max(prices[i]-prices[i-1], dynamic.maxProfit)
		if max > dynamic.maxProfit {
			dynamic.maxProfit = max
			dynamic.minIdx = i - 1
		}
		if prices[i] < prices[dynamic.minIdx] {
			dynamic.minIdx = i
		}

	}

	return dynamic.maxProfit
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
