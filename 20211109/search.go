package _0211109

/**
 * 统计一个数字在排序数组中出现的次数。：https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
 */

func search(nums []int, target int) int {
	var c int

	if len(nums) == 0 {
		return c
	}

	l := 0
	r := len(nums)
	m := 0

	for l <= r {
		m = (r + l) / 2

		if m >= len(nums) || nums[m] == target {
			break
		}

		if nums[m] > target {
			r = m - 1
		}

		if nums[m] < target {
			l = m + 1
		}
	}

	if m >= len(nums) || nums[m] != target {
		return c
	}

	for i := m; i >= 0; i-- {
		if nums[i] != target {
			break
		}
		c++
	}

	for j := m + 1; j < len(nums); j++ {
		if nums[j] != target {
			break
		}
		c++
	}

	return c
}
