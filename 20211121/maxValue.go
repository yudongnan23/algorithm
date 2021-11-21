package _0211121

/**
 * 礼物的最大价值：https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/
 */

func maxValue(g [][]int) int {
	if len(g) == 0 || len(g[0]) == 0 {
		return 0
	}

	d := make([][]int, len(g))

	for i := 0; i < len(g); i++ {
		d[i] = make([]int, len(g[i]))
		for j := 0; j < len(g[i]); j++ {
			if i == 0 && j == 0 {
				d[i][j] = g[i][j]
				continue
			}
			if i == 0 {
				d[i][j] = d[i][j-1] + g[i][j]
				continue
			}
			if j == 0 {
				d[i][j] = d[i-1][j] + g[i][j]
				continue
			}

			d[i][j] = max(d[i][j-1], d[i-1][j]) + g[i][j]
		}
	}

	return d[len(d)-1][len(d[len(d)-1])-1]
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}