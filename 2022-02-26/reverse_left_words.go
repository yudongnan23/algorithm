package _022_02_26

// 左旋字符串
// leetcode链接：https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/

/**
 * 思路：字符串切片
 */

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}
