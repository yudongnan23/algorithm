package _022_03_06

// 第一个只出现一次的字符
// leetcode链接：https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/

/**
 * 思路：遍历两遍
 */

func firstUniqChar(s string) byte {
	cnt := [26]int{}

	for _, ch := range s {
		cnt[ch-'a'] += 1
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
