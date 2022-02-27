package _022_02_27

// 找出数组中重复的数字
// leetcode链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

/**
 * 思路：遍历数组，并使用map记录出现的数字
 */

func findRepeatNumber(nums []int) int {
	dumpMapping := make(map[int]bool, len(nums))

	for i := 0; i < len(nums); i++ {
		if _, ok := dumpMapping[nums[i]]; ok {
			return nums[i]
		}

		dumpMapping[nums[i]] = true
	}

	return 0
}
