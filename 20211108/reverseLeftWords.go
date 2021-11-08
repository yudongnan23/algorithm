package _0211108

/**
 * 左旋转字符串
 */

func reverseLeftWords(s string, n int) string {
	if n == 0 || n >= len(s) {
		return s
	}

	return s[n:] + s[0:n]
}
