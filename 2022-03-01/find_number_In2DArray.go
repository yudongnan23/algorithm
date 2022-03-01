package _022_03_01

// 二维数组中的查找
// leetcode链接：https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/

/**
 * 思路：遍历二维数组，并对每个数组进行二分查找
 */

func findNumberIn2DArray(matrix [][]int, target int) bool {
	var exist bool

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return exist
	}

	for i := range matrix {
		if matrix[i][0] > target {
			continue
		}
		if matrix[i][len(matrix[i])-1] < target {
			continue
		}

		exist = binarySearch(matrix[i], target)
		if exist {
			break
		}
	}
	return exist
}

func binarySearch(nums []int, target int) bool {
	var exist bool

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			exist = true
			break
		}

		if nums[mid] < target {
			left = mid + 1
		}

		if nums[mid] > target {
			right = mid - 1
		}
	}

	return exist
}
