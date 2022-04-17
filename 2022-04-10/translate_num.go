package _022_04_10

import "strconv"

// 把数字翻译成字符串
// leetcode连接：htt:ps"//leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/

/**
 * 思路：动态规划
 */

var dict = map[string]bool{
	"10": true,
	"11": true,
	"12": true,
	"13": true,
	"14": true,
	"15": true,
	"16": true,
	"17": true,
	"18": true,
	"19": true,
	"20": true,
	"21": true,
	"22": true,
	"23": true,
	"24": true,
	"25": true,
}

func translateNum(num int) int {
	ss := strconv.Itoa(num)
	if len(ss) == 1 {
		return 1
	}

	d := make([]int, len(ss))
	d[0] = 1
	if dict[ss[0:2]] {
		d[1] = 2
	} else {
		d[1] = 1
	}

	for i := 2; i < len(ss); i++ {
		if dict[ss[i-1:i+1]] {
			d[i] = d[i-1] + d[i-2]
			continue
		}

		d[i] = d[i-1]
	}
	return d[len(ss)-1]
}
