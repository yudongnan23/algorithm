package main

/**
 * 无重复字符的最长子串: https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
 */

func lengthOfLongestSubstring(s string) int {
	m := make(map[string]bool, 0)
	count := 0
	start := 0
	end := 1
	for i := range s {
		if _, ok := m[string(s[i])]; ok {
			m = make(map[string]bool, 0)
			m[string(s[i])] = true
			if start - end > count {
				count = start - end
			}
			start = i
			end = i + 1
			continue
		}

		m[string(s[i])] = true
		end = end + 1
		if end - start > count {
			count = end - start
		}
	}
	return count
}

