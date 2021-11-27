package _0211127

/**
 * 和为s的两个数字：https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
 */

// 双指针
func twoSum2(nn []int, t int) []int {
	rlt := make([]int, 2)
	if len(nn) <= 1 {
		return rlt
	}

	l := 0
	r := len(nn) - 1

	for l < r {
		s := nn[l] + nn[r]
		if s == t {
			rlt[0] = nn[l]
			rlt[1] = nn[r]
			break
		}
		if s > t {
			r--
		}

		if s < t {
			l++
		}
	}
	return rlt

}

// 暴力破解
func twoSum(nn []int, t int) []int {
	dump := make(map[int]bool, len(nn))

	r := make([]int, 2)
	for i := 0; i < len(nn); i++ {
		if dump[t-nn[i]] {
			r[0] = nn[i]
			r[1] = t - nn[i]
			break
		}
		dump[nn[i]] = true
	}
	return r
}
