package _022_03_01

// 旋转数组中的最小数字
// leetcode链接：https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/

/**
 * 思路：使用两个指针，一头一尾进行双向遍历
 */

func minArray(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left] >= nums[right] {
			left++
			continue
		}

		if nums[left] < nums[right] {
			right--
		}
	}
	return nums[left]
}
