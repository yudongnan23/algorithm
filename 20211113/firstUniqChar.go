package _0211113

/**
 * 第一个只出现一次的字符: https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
 */

import "strings"

func firstUniqChar(s string) byte {
	var first byte

	for i := 0; i < len(s); i++ {
		if strings.IndexByte(s[0:i], s[i]) == -1 && strings.IndexByte(s[i+1:], s[i]) == -1 {
			first = s[i]
			break
		}

	}

	if first == 0 {
		return " "[0]
	}

	return first
}
