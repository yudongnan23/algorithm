package _0211116

/**
 * 斐波那契数列：https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
 */

import (
	"math/big"
	"strconv"
)

func fib(n int) int {
	d := make([]*big.Int, 2)
	d[0] = big.NewInt(0)
	d[1] = big.NewInt(1)

	for i := 2; i <= n; i++ {
		s := big.NewInt(0)
		d = append(d, s.Add(d[i-1], d[i-2]))
	}

	s := big.NewInt(0)
	i, _ := strconv.Atoi(s.Mod(d[n], big.NewInt(1000000007)).String())
	return i
}
