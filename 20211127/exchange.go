package _0211127

/**
 * 调整数组顺序使奇数位于偶数前面：https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
 */

func exchange(nn []int) []int {
	if len(nn) <= 1 {
		return nn
	}

	l := 0
	r := 1

	for r < len(nn) {
		if !isOdd(nn[l]) && isOdd(nn[r]) {
			t := nn[r]
			nn[r] = nn[l]
			nn[l] = t
		}

		if isOdd(nn[l]) {
			l++
		}

		r++
	}
	return nn
}

func isOdd(n int) bool {
	return n%2 == 1
}
