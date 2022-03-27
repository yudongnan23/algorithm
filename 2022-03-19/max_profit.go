package _022_03_19

// 股票最大利润
// leetcode链接：https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/

/**
 * 思路：动态规划
 */

type profitElem struct {
	maxProfit int
	minIdx    int
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	d := profitElem{
		maxProfit: 0,
		minIdx:    0,
	}

	for i := 1; i < len(prices); i++ {
		maxSum := prices[i] - prices[d.minIdx]
		if maxSum > d.maxProfit {
			d.maxProfit = maxSum
		}
		if prices[i] < prices[d.minIdx] {
			d.minIdx = i
		}
	}
	return d.maxProfit
}
