package _022_03_26

// 连续子数组的最大和
// leetcode链接：https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/

/**
 * 思路：动态规划
 */

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dynamic := make([]int, len(nums))
	dynamic[0] = nums[0]

	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if sum := dynamic[i-1] + nums[i]; sum > dynamic[i-1] {
			dynamic[i] = sum
		} else {
			dynamic[i] = nums[i]
		}

		if dynamic[i] > max {
			max = dynamic[i]
		}
	}

	return max
}
