package _022_02_27

// 统计重复数字在排序数组中出现的次数
// leetcode链接：https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/

/**
 * 思路：二分查找
 */

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := len(nums) - 1
	startIndex := 0

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			startIndex = mid
			break
		}

		if nums[mid] > target {
			right = mid - 1
		}

		if nums[mid] < target {
			left = mid + 1
		}
	}

	if nums[startIndex] == target {
		return countLeft(&nums, startIndex, target) + countRight(&nums, startIndex, target) + 1
	}

	return 0
}

func countLeft(nums *[]int, startIndex, target int) int {
	var count int

	for i := startIndex - 1; i >= 0; i-- {
		if (*nums)[i] != target {
			break
		}
		count++
	}

	return count
}

func countRight(nums *[]int, startIndex, target int) int {
	var count int

	for i := startIndex + 1; i < len(*(nums)); i++ {
		if (*nums)[i] != target {
			break
		}
		count++
	}

	return count
}
