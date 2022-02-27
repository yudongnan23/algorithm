package _022_02_27

// 寻找0 ~ n-1中缺失的数字
// leetcode链接：https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/

/**
 * 思路：等差数列求和
 */

func missingNumber(nums []int) int {
	return arithmeticSum(len(nums)) - arraySum(nums)
}

func arithmeticSum(n int) int {
	return ((n - 1) * n) / 2
}

func arraySum(nums []int) int {
	var sum int

	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}

	return sum
}

func missingNumberOptimize(nums []int) int {
	nums = append(nums, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}

	return -1
}
