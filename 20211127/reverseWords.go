package _0211127

import "strings"

/**
 * 翻转单次的顺序：https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/
 */

func reverseWords(s string) string {
	ss := strings.Fields(s)
	l, r := 0, len(ss)-1
	for l < r {
		ss[l], ss[r] = ss[r], ss[l]
		l++
		r--
	}

	return strings.Join(ss, " ")
}
