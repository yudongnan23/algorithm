package main

import "strings"

/**
 * 无重复字符的最长子串: https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
 * acacacacacacacacac
 * sssd
 * dvdf
 * "tmmzuxt"
 */
func lengthOfLongestSubstring(s string) int {
	placeMapping := make(map[string]int, len(s))
	maxLength := 0
	start := 0
	end := 0

	for i := range s {
		end = i

		if strings.Contains(s[start:end], string(s[i])) {
			if end-start > maxLength {
				maxLength = end - start
			}

			start = placeMapping[string(s[i])] + 1
		}
		placeMapping[string(s[i])] = i
	}

	if end-start+1 > maxLength {
		maxLength = end - start + 1
	}

	return maxLength
}

/**
 * 暴力破解
 */
func lengthOfLongestSubstringMethod2(s string) int {
	maxLength := 0
	for i := range s {
		ss := s[i:]
		if len(ss) < maxLength {
			continue
		}
		curCount := maxLengthContinuous(ss)
		if curCount > maxLength {
			maxLength = curCount
		}
	}
	return maxLength
}

func maxLengthContinuous(ss string) int {
	dumMapping := make(map[string]bool)
	count := 0
	for _, s := range ss {
		if _, ok := dumMapping[string(s)]; ok {
			break
		}
		count = count + 1
		dumMapping[string(s)] = true
	}
	return count
}
