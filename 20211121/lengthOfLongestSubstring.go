package _0211121

/**
 * 最长不包含重复字符的子字符串：https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
 */

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := 1
	start := 0
	end := 1
	mapping := make(map[string]int, 0)
	mapping[string(s[0])] = 0

	for i := 1; i < len(s); i++ {
		if _, ok := mapping[string(s[i])]; ok && mapping[string(s[i])] >= start {
			start = mapping[string(s[i])]
			mapping[string(s[i])] = i
		}
		end++
		if end-start > max {
			max = end - start
		}
	}
	return max
}
