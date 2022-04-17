package _022_04_17

// 最长不含重复字符的子字符串
// leetcode链接：https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/

/**
 * 思路：动态规划
 */

const _move = 1

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	dump := make(map[byte]int, len(s))
	lastMinIndex := -1
	var max int

	for i := range s {
		if dump[s[i]] > 0 && dump[s[i]]-_move >= lastMinIndex {
			lastMinIndex = dump[s[i]] - _move
			dump[s[i]] = i + _move
			continue
		}

		if i-lastMinIndex > max {
			max = i - lastMinIndex
		}
		dump[s[i]] = i + _move
	}

	return max
}
