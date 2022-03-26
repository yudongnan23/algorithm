package _022_03_26

// 连续子数组的最大和
// leetcode链接：https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/

/**
 * 思路：
 */

type maxSub struct {
	maxSum      int
	intervalSum int
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSub := maxSub{
		maxSum:      nums[0],
		intervalSum: 0,
	}

	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxSub.intervalSum = maxSub.intervalSum + nums[i]
			continue
		}
		if nums[i] > maxSub.maxSum {
			maxSub.maxSum = nums[i]
			maxSub.intervalSum = 0
			if nums[i] > max {
				max = nums[i]
			}
			continue
		}
		if curSum := nums[i] + maxSub.maxSum + maxSub.intervalSum; curSum > maxSub.maxSum {
			maxSub.maxSum = curSum
			maxSub.intervalSum = 0
			if curSum > max {
				max = curSum
			}
			continue
		}
	}

	return max
}
