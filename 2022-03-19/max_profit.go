package _022_03_19

// 股票最大利润
// leetcode链接：https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/

/**
 * 思路：动态规划
 */

func maxProfit(prices []int) int {
	if len(prices) <= 1{
		return 0
	}

	var maxProfit int
	var minIndex int

	for i := 1; i < len(prices); i++ {
		if prices[i] - prices[minIndex] > maxProfit {
			maxProfit = prices[i] - prices[minIndex]
		}

		if prices[i] < prices[minIndex] {
			minIndex = i
		}
	}

	return maxProfit
}
