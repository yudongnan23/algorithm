package main

/**
 * 无重复字符的最长子串: https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
 * acacacacacacacacac
 * sssd
 * dvdf
 * "tmmzuxt"
 */

func lengthOfLongestSubstring(s string) int {
	m := make(map[string]int, 0)
	count := 0
	start := -1
	for end := range s {
		if _, ok := m[string(s[end])]; ok {
			start = m[string(s[end])]
			m = deletePre(m, s, m[string(s[end])])
			m[string(s[end])] = end
			continue
		}

		if end-start > count {
			count = end - start
		}
		m[string(s[end])] = end
	}
	return count
}

func deletePre(m map[string]int, s string, end int)map[string]int{
	for i := 0; i <= end; i ++{
		delete(m, string(s[i]))
	}
	return m
}