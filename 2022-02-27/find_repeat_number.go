package _022_02_27

// 找出数组中重复的数字
// leetcode链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

/**
 * 思路:
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
