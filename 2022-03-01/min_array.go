package _022_03_01

// 旋转数组中的最小数字
// leetcode链接：https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/

/**
 * 思路：
 */

func minArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] <= nums[right] {
			right = mid - 1
		}

		if nums[mid] >= nums[right] {
			left = mid + 1
		}
	}

	if nums[left] < nums[right] {
		return nums[left]
	}
	return nums[right]
}

