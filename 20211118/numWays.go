package _0211118

import (
	"math/big"
	"strconv"
)

/**
 * 青蛙跳台阶： https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
 * f(n) = f(n-1) + f(n-2)
 */

func numWays(n int) int {
	w := make([]*big.Int, 3)
	if n > 2 {
		w = make([]*big.Int, n+1)
	}
	w[0] = big.NewInt(1)
	w[1] = big.NewInt(1)
	w[2] = big.NewInt(2)

	for i := 3; i <= n; i++ {
		s := big.NewInt(0)
		w[i] = s.Add(w[i-1], w[i-2])
	}


	r, _ := strconv.Atoi(big.NewInt(0).Mod(w[n], big.NewInt(1000000007)).String())
	return r
}
