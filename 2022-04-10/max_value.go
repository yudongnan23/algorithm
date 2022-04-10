package _022_04_10

// 礼物的最大价值
// leetcode连接：https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/

/**
 * 思路：动态规划
 */

func maxValue(gird [][]int) int {
	if len(gird) == 0 || len(gird[0]) == 0 {
		return 0
	}

	for i := 0; i < len(gird); i++ {
		for j := 0; j < len(gird[i]); j++ {
			var up, left, curSum int

			if i-1 >= 0 {
				up = gird[i-1][j]
			}

			if j-1 >= 0 {
				left = gird[i][j-1]
			}

			curSum = left + gird[i][j]
			if up+gird[i][j] > left+gird[i][j] {
				curSum = up + gird[i][j]
			}

			gird[i][j] = curSum
		}
	}

	return gird[len(gird)-1][len(gird[0])-1]
}
