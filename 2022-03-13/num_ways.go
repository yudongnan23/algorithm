package _022_03_13

import big2 "math/big"

// 🐸跳台阶
// leetcode链接：https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/

/**
 * 思路：
 */

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}

	dynamic := make([]*big2.Int, n+1)
	dynamic[0] = big2.NewInt(1)
	dynamic[1] = big2.NewInt(1)

	for i := 2; i < n+1; i++ {
		big := big2.NewInt(0)
		dynamic[i] = big.Add(dynamic[i-1], dynamic[i-2])
	}

	big := big2.NewInt(0)
	return int(big.Mod(dynamic[n], big2.NewInt(_mod)).Int64())
}
